package infra

import (
	"context"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewTraceProviderBuilder(name string) *traceProvider {
	return &traceProvider{
		name: name,
	}
}

type traceProvider struct {
	name     string
	exporter trace.SpanExporter
}

func (t *traceProvider) SetExporter(exp trace.SpanExporter) *traceProvider {
	t.exporter = exp
	return t
}

func (t *traceProvider) Build() (*trace.TracerProvider, gcommon.CloseFn, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(t.name),
		),
	)
	if err != nil {
		return nil, nil, err
	}

	bsp := trace.NewBatchSpanProcessor(t.exporter)

	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	return tracerProvider, func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := t.exporter.Shutdown(ctx); err != nil {
			return err
		}
		return err
	}, nil
}
