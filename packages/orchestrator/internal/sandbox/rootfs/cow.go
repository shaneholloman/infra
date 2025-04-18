package rootfs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/bits-and-blooms/bitset"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/block"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/nbd"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/template"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
	"github.com/e2b-dev/infra/packages/shared/pkg/utils"
)

type CowDevice struct {
	overlay *block.Overlay
	mnt     *nbd.DirectPathMount

	ready *utils.SetOnce[string]

	blockSize   int64
	BaseBuildId string

	finishedOperations chan struct{}
	devicePool         *nbd.DevicePool

	tracer trace.Tracer
}

func NewCowDevice(tracer trace.Tracer, rootfs *template.Storage, cachePath string, blockSize int64, devicePool *nbd.DevicePool) (*CowDevice, error) {
	size, err := rootfs.Size()
	if err != nil {
		return nil, fmt.Errorf("error getting device size: %w", err)
	}

	cache, err := block.NewCache(size, blockSize, cachePath, false)
	if err != nil {
		return nil, fmt.Errorf("error creating cache: %w", err)
	}

	overlay := block.NewOverlay(rootfs, cache, blockSize)

	mnt := nbd.NewDirectPathMount(tracer, overlay, devicePool)

	return &CowDevice{
		tracer:             tracer,
		mnt:                mnt,
		overlay:            overlay,
		ready:              utils.NewSetOnce[string](),
		blockSize:          blockSize,
		finishedOperations: make(chan struct{}, 1),
		BaseBuildId:        rootfs.Header().Metadata.BaseBuildId.String(),
		devicePool:         devicePool,
	}, nil
}

func (o *CowDevice) Start(ctx context.Context) error {
	deviceIndex, err := o.mnt.Open(ctx)
	if err != nil {
		return o.ready.SetError(fmt.Errorf("error opening overlay file: %w", err))
	}

	return o.ready.SetValue(nbd.GetDevicePath(deviceIndex))
}

func (o *CowDevice) Export(parentCtx context.Context, out io.Writer, stopSandbox func(ctx context.Context) error) (*bitset.BitSet, error) {
	childCtx, childSpan := o.tracer.Start(parentCtx, "cow-export")
	defer childSpan.End()

	cache, err := o.overlay.EjectCache()
	if err != nil {
		return nil, fmt.Errorf("error ejecting cache: %w", err)
	}

	// the error is already logged in go routine in SandboxCreate handler
	go stopSandbox(childCtx)

	select {
	case <-o.finishedOperations:
		break
	case <-childCtx.Done():
		return nil, fmt.Errorf("timeout waiting for overlay device to be released")
	}
	telemetry.ReportEvent(childCtx, "sandbox stopped")

	dirty, err := cache.Export(out)
	if err != nil {
		return nil, fmt.Errorf("error exporting cache: %w", err)
	}

	telemetry.ReportEvent(childCtx, "cache exported")

	err = cache.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing cache: %w", err)
	}

	return dirty, nil
}

func (o *CowDevice) Close(ctx context.Context) error {
	childCtx, childSpan := o.tracer.Start(ctx, "cow-close")
	defer childSpan.End()

	var errs []error

	err := o.mnt.Close(childCtx)
	if err != nil {
		errs = append(errs, fmt.Errorf("error closing overlay mount: %w", err))
	}

	o.finishedOperations <- struct{}{}

	err = o.overlay.Close()
	if err != nil {
		errs = append(errs, fmt.Errorf("error closing overlay cache: %w", err))
	}

	devicePath, err := o.ready.Wait()
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting overlay path: %w", err))

		return errors.Join(errs...)
	}

	slot, err := nbd.GetDeviceSlot(devicePath)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting overlay slot: %w", err))

		return errors.Join(errs...)
	}

	attempts := 0
	for {
		attempts++
		err := o.devicePool.ReleaseDevice(slot)
		if errors.Is(err, nbd.ErrDeviceInUse{}) {
			if attempts%100 == 0 {
				zap.L().Info("error releasing overlay device", zap.Int("attempts", attempts), zap.Error(err))
			}

			time.Sleep(500 * time.Millisecond)

			continue
		}

		if err != nil {
			return fmt.Errorf("error releasing overlay device: %w", err)
		}

		break
	}

	zap.L().Info("overlay device released")

	return nil
}

func (o *CowDevice) Path() (string, error) {
	return o.ready.Wait()
}
