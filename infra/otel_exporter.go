package infra

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
)

func NewOTLP(endpoint string) *otlptrace.Exporter {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)

	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create the collector trace exporter")
	}

	return traceExp
}
