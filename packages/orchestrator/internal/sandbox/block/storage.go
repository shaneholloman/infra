package block

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/build"
	"github.com/e2b-dev/infra/packages/shared/pkg/storage/header"
)

type Storage struct {
	header *header.Header
	source *chunker
}

func NewStorage(
	ctx context.Context,
	store *build.Store,
	buildId string,
	storeKeySuffix string,
	blockSize int64,
	cachePath string,
	isSnapshot bool,
) (*Storage, error) {
	id, err := uuid.Parse(buildId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse build id: %w", err)
	}

	var h *header.Header

	if isSnapshot {
		headerObject, err := store.Get(id.String() + "/" + storeKeySuffix + ".header")
		if err != nil {
			return nil, fmt.Errorf("failed to get header object: %w", err)
		}

		diffHeader, err := header.Deserialize(headerObject)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize header: %w", err)
		}

		h = diffHeader
	} else {
		object, err := store.Get(id.String() + "/" + storeKeySuffix)
		if err != nil {
			return nil, fmt.Errorf("failed to get object: %w", err)
		}

		size, err := object.Size()
		if err != nil {
			return nil, fmt.Errorf("failed to get object size: %w", err)
		}

		h = header.NewHeader(&header.Metadata{
			BuildId:     id,
			BaseBuildId: id,
			Size:        uint64(size),
			Version:     1,
			BlockSize:   uint64(blockSize),
			Generation:  1,
		}, nil)
	}

	b := build.NewFromStorage(h, store, storeKeySuffix)

	chunker, err := newChunker(ctx, int64(h.Metadata.Size), blockSize, b, cachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create chunker: %w", err)
	}

	return &Storage{
		source: chunker,
		header: h,
	}, nil
}

func (d *Storage) ReadAt(p []byte, off int64) (int, error) {
	return d.source.ReadAt(p, off)
}

func (d *Storage) Size() (int64, error) {
	return int64(d.header.Metadata.Size), nil
}

func (d *Storage) Close() error {
	return d.source.Close()
}

func (d *Storage) Slice(off, length int64) ([]byte, error) {
	return d.source.Slice(off, length)
}

func (d *Storage) Header() *header.Header {
	return d.header
}
