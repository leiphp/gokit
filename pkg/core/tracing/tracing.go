package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var tracer = otel.Tracer("gokit")

func Init(serviceName string) error {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return tracer.Start(ctx, name)
}
