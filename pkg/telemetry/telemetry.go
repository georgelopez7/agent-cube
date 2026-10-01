package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.36.0"
	adktelemetry "google.golang.org/adk/telemetry"
)

// New initializes ADK OpenTelemetry providers and registers them globally.
func New(ctx context.Context, serviceName string) (*adktelemetry.Providers, error) {
	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceNameKey.String(serviceName),
		attribute.String("openinference.project.name", "agentcube"),
	))
	if err != nil {
		return nil, err
	}

	providers, err := adktelemetry.New(ctx, adktelemetry.WithResource(res))
	if err != nil {
		return nil, err
	}

	providers.SetGlobalOtelProviders()
	return providers, nil
}
