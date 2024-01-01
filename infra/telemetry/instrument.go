package telemetry

import (
	"go.opentelemetry.io/otel"
)

var Trace = otel.Tracer("realworld-go")
