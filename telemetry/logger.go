package telemetry

import (
	"context"
	"errors"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"os"
	"time"
)

func NewLogger(ctx context.Context, provider *log.LoggerProvider, name string) (*zap.Logger, func() error, error) {
	// Initialize a zap logger with the otelzap bridge core.
	// This method actually doesn't log anything on your STDOUT, as everything
	// is shipped to a configured otel endpoint.
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		otelzap.NewCore(name, otelzap.WithLoggerProvider(provider)),
	)
	logger := zap.New(core)
	return logger, func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		return errors.Join(
			logger.Sync(),
			provider.Shutdown(ctx),
		)
	}, nil
}

func NewLoggerProvider(ctx context.Context, res *resource.Resource, endpoint string) (*log.LoggerProvider, error) {
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(endpoint),
		otlploggrpc.WithInsecure(), // remove if using TLS
		otlploggrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		return nil, err
	}
	processor := log.NewBatchProcessor(exporter)
	provider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(processor),
	)
	return provider, nil
}
