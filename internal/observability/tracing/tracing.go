package tracing

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func NewTracer(serviceName string) func(context.Context) error {
	ctx := context.Background()

	otelEndpoint := os.Getenv("OTLP_ENDPOINT")
	if otelEndpoint == "" {
		otelEndpoint = "localhost:4318"
	}

	// EXPORTER (TO OPENTELEMETRY USING HTTP)
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(otelEndpoint), otlptracehttp.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create opentelemetry exporter: %v", err)
	}

	// IMPORTANT: PROPAGATE TRACE CONTEXT
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// TRACER PROVIDER
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName), // IMPORTANT: SET SERVICE NAME
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
