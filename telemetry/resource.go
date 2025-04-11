package telemetry

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func Resource(values ...attribute.KeyValue) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			values...,
		))
}

func DefaultResource() *resource.Resource {
	return resource.Default()
}

func defaultResource(serviceName string) *resource.Resource {
	if serviceName == "" {
		return DefaultResource()
	} else {
		r, err := Resource(semconv.ServiceName(serviceName))
		if err != nil {
			return DefaultResource()
		}
		return r
	}
}
