package logs

import (
	"context"
	"io"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"

	"github.com/e2b-dev/infra/packages/shared/pkg/consts"
	"github.com/e2b-dev/infra/packages/shared/pkg/logs/exporter"
)

const (
	OrchestratorServiceName = "orchestrator"
	cpuUsageThreshold       = 0.85
	memoryUsageThreshold    = 0.85
)

type SandboxLogExporter struct {
	logger *zerolog.Logger
}

func newSandboxLogExporter(serviceName string) *SandboxLogExporter {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	ctx := context.Background()
	exporters := []io.Writer{exporter.NewHTTPLogsExporter(ctx, consts.LogsProxyAddress)}

	l := zerolog.
		New(io.MultiWriter(exporters...)).
		With().
		Timestamp().
		Logger().
		Level(zerolog.DebugLevel).
		With().Str("logger", serviceName).Logger()

	return &SandboxLogExporter{
		logger: &l,
	}
}

var sandboxLogExporter *SandboxLogExporter
var sandboxLogExporterMU = sync.Mutex{}

func getSandboxLogExporter() *SandboxLogExporter {
	sandboxLogExporterMU.Lock()
	defer sandboxLogExporterMU.Unlock()

	if sandboxLogExporter == nil {
		sandboxLogExporter = newSandboxLogExporter(OrchestratorServiceName)
	}

	return sandboxLogExporter
}

type SandboxLogger struct {
	exporter              *SandboxLogExporter
	internal              bool
	instanceID            string
	envID                 string
	teamID                string
	cpuMax                int32
	cpuWasAboveTreshold   atomic.Bool
	memoryMBMax           int32
	memoryWasAbove        atomic.Int32
	healthCheckWasFailing atomic.Bool
}

func NewSandboxLogger(
	instanceID string,
	envID string,
	teamID string,
	cpuMax int32,
	memoryMax int32,
	internal bool,
) *SandboxLogger {
	logsExporter := getSandboxLogExporter()
	return &SandboxLogger{
		exporter:    logsExporter,
		instanceID:  instanceID,
		internal:    internal,
		envID:       envID,
		teamID:      teamID,
		cpuMax:      cpuMax,
		memoryMBMax: memoryMax,
	}
}

func (l *SandboxLogger) sendEvent(logger *zerolog.Event, format string, v ...interface{}) {
	logger.
		Str("instanceID", l.instanceID).
		Str("envID", l.envID).
		Str("teamID", l.teamID).
		Bool("internal", l.internal).
		Msgf(format, v...)
}

func (l *SandboxLogger) GetInternalLogger() *SandboxLogger {
	if l.internal {
		return l
	}

	return NewSandboxLogger(l.instanceID, l.envID, l.teamID, l.cpuMax, l.memoryMBMax, true)
}

func (l *SandboxLogger) Errorf(
	format string,
	v ...interface{},
) {
	l.sendEvent(l.exporter.logger.Error(), format, v...)
}

func (l *SandboxLogger) Warnf(
	format string,
	v ...interface{},
) {
	l.sendEvent(l.exporter.logger.Warn(), format, v...)
}

func (l *SandboxLogger) Infof(
	format string,
	v ...interface{},
) {
	l.sendEvent(l.exporter.logger.Info(), format, v...)
}
func (l *SandboxLogger) Debugf(
	format string,
	v ...interface{},
) {
	l.sendEvent(l.exporter.logger.Debug(), format, v...)
}

func (l *SandboxLogger) CPUUsage(cpu float64) {
	// Round to 3 decimal places and cap at cpuMax
	cpu = math.Min(float64(int(cpu*1000))/1000, float64(l.cpuMax))
	if cpu > cpuUsageThreshold*float64(l.cpuMax) {
		l.cpuWasAboveTreshold.Store(true)

		l.exporter.logger.Warn().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Float64("cpu", cpu).
			Int32("cpuMax", l.cpuMax).
			Msgf("cpu usage reached %d %% of total cpu", int(cpu/float64(l.cpuMax)*100))
	} else if l.cpuWasAboveTreshold.Load() && cpu <= cpuUsageThreshold*float64(l.cpuMax) {
		l.cpuWasAboveTreshold.Store(false)
		l.exporter.logger.Warn().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Float64("cpu", cpu).
			Int32("cpuMax", l.cpuMax).
			Msgf("cpu usage fell below %d %% of total cpu", int(cpuUsageThreshold*100))
	}
}

func (l *SandboxLogger) MemoryUsage(memoryMB float64) {
	// Cap at memoryMBMax
	memoryMB = math.Min(memoryMB, float64(l.memoryMBMax))
	if memoryMB > memoryUsageThreshold*float64(l.memoryMBMax) && int32(memoryMB) > l.memoryWasAbove.Load() {
		l.memoryWasAbove.Store(int32(memoryMB))
		l.exporter.logger.Warn().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Float64("memoryMB", memoryMB).
			Int32("memoryMBMax", l.memoryMBMax).
			Msgf("memoryMB usage reached %d %% of memoryMB", int(memoryMB/float64(l.memoryMBMax)*100))
		return
	}
}

func (l *SandboxLogger) Healthcheck(ok bool, alwaysReport bool) {
	if !ok && !l.healthCheckWasFailing.Load() {
		l.healthCheckWasFailing.Store(true)

		l.exporter.logger.Error().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Bool("healthcheck", ok).
			Msg("Sandbox healthcheck started failing")
		return
	}
	if ok && l.healthCheckWasFailing.Load() {
		l.healthCheckWasFailing.Store(false)

		l.exporter.logger.Warn().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Bool("healthcheck", ok).
			Msg("Sandbox healthcheck recovered")

		return
	}

	if alwaysReport {
		var msg string
		if ok {
			msg = "Control sandbox healthcheck was successful"
		} else {
			msg = "Control sandbox healthcheck failed"
		}

		l.exporter.logger.Info().
			Str("instanceID", l.instanceID).
			Str("envID", l.envID).
			Str("teamID", l.teamID).
			Bool("healthcheck", ok).
			Msg(msg)
	}

}
