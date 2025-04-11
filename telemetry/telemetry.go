package telemetry

import (
	"context"
	"errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/metric"
	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	tracerSdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"io"
	"os"
	"time"
)

type Telemetry interface {
	Metrics() metric.MeterProvider
	Traces() trace.TracerProvider
	io.Closer
}

type abcTelemetry struct {
	traces  *tracerSdk.TracerProvider
	metrics *metricSdk.MeterProvider
	Logger  *zap.Logger

	_sync func() error
}

func (t *abcTelemetry) Metrics() metric.MeterProvider {
	return t.metrics
}

func (t *abcTelemetry) Traces() trace.TracerProvider {
	return t.traces
}

func (t *abcTelemetry) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return errors.Join(
		t._sync(),
		t.metrics.Shutdown(ctx),
		t.traces.Shutdown(ctx),
	)
}

// Init serviceName is overriden by env var OTEL_SERVICE_NAME
func Init(ctx context.Context, serviceName string) (*abcTelemetry, error) {
	_serviceName := os.Getenv(consts.ServiceNameEnvVarName)
	if _serviceName != "" {
		serviceName = _serviceName
	}
	endpoint := os.Getenv(consts.OtelEndpointEnvVarName)
	if endpoint == "" {
		endpoint = consts.DefaultTelemetryEndpoint
	}
	resource := defaultResource(serviceName)

	var t abcTelemetry
	// Use a working LoggerProvider implementation instead e.g. use go.opentelemetry.io/otel/sdk/log.
	logs, err := NewLoggerProvider(ctx, resource, endpoint)
	if err != nil {
		return &t, err
	}
	global.SetLoggerProvider(logs)

	logger, sync, err := NewLogger(ctx, logs, serviceName)
	if err != nil {
		return &t, err
	}
	t._sync = sync
	t.Logger = logger
	zap.ReplaceGlobals(logger)

	metrics, err := NewMeterProvider(ctx, resource, endpoint)
	if err != nil {
		return &t, err
	}
	t.metrics = metrics
	otel.SetMeterProvider(metrics)

	traces, err := NewTracerProvider(ctx, resource, endpoint)
	if err != nil {
		return &t, err
	}
	t.traces = traces
	otel.SetTracerProvider(traces)

	return &t, nil
}
