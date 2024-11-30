package oteltracer

import (
	"context"
	"fmt"
	"strings"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract/tracercontract"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"
)

var _ tracercontract.Tracer = (*OpenTelemetry)(nil)

// OpenTelemetry encapsulates the OpenTelemetry components required for distributed tracing.
// It provides a tracer instance for span creation and a tracer provider for managing lifecycle and configurations.
// This struct facilitates the seamless integration of OpenTelemetry into applications.
type OpenTelemetry struct {
	cfg            Config
	tracer         trace.Tracer             // Tracer instance for creating spans.
	tracerProvider *sdktrace.TracerProvider // TracerProvider to manage tracer lifecycle and export spans.
}

// New creates a new instance of OpenTelemetry with the provided configuration.
//
// Parameters:
// - cfg: The configuration containing settings for OpenTelemetry, including service name, OTLP endpoint, and
// other options.
//
// Returns:
// - *OpenTelemetry: An unconfigured OpenTelemetry instance, ready for setup.
func New(cfg Config) *OpenTelemetry {
	return &OpenTelemetry{
		cfg:            cfg,
		tracer:         nil,
		tracerProvider: nil,
	}
}

// Configure initializes the OpenTelemetry by setting up the tracer provider and exporter.
//
// This method validates the configuration, creates an OTLP exporter, and configures the tracer provider
// with the appropriate settings. The returned instance is fully functional for tracing operations.
//
// Parameters: None.
//
// Returns:
// - error: An error if the configuration or initialization fails, such as missing endpoints or exporter setup issues.
func (o *OpenTelemetry) Configure() error {
	if len(o.cfg.Endpoint) == 0 {
		return ErrOTLPEndpointRequired
	}

	exp, err := createExporter(o.cfg)
	if err != nil {
		return err
	}

	tp := createTraceProvider(o.cfg, exp)

	o.tracer = tp.Tracer(o.cfg.AppName)
	o.tracerProvider = tp

	return nil
}

// Span creates and starts a new span with the specified name and attributes.
//
// Parameters:
// - ctx: The context to associate with the span.
// - name: The name of the span.
//
// Returns:
// - context.Context: The context associated with the span.
// - trace.Span: The created span instance.
func (o *OpenTelemetry) Span(
	ctx context.Context,
	name string,
) (context.Context, tracercontract.Span) {
	ctx, span := o.tracer.Start(ctx, name) //nolint:spancheck // nothing

	return ctx, NewSpan(span) //nolint:spancheck // nothing
}

// Shutdown gracefully shuts down the tracer provider and flushes any remaining spans to the exporter.
//
// Parameters:
// - ctx: The context for managing the shutdown operation.
//
// Returns:
// - error: An error if the shutdown operation fails.
func (o *OpenTelemetry) Shutdown(ctx context.Context) error {
	if o.tracer == nil {
		return nil // No tracer initialized, no need to shut down
	}

	if o.tracerProvider == nil {
		return nil // No trace provider initialized, nothing to shut down
	}

	if err := o.tracerProvider.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down tracer provider: %w", err)
	}

	return nil
}

// createExporter initializes an OpenTelemetry SpanExporter based on the provided configuration.
// It supports both HTTP and gRPC exporters, depending on the specified endpoint.
//
// Parameters:
// - cfg: Configuration containing the OTLP endpoint details.
//
// Returns:
// - sdktrace.SpanExporter: The initialized span exporter.
// - error: An error if exporter creation fails.
func createExporter(cfg Config) (*otlptrace.Exporter, error) {
	var (
		// exp sdktrace.SpanExporter
		exp *otlptrace.Exporter
		err error
	)

	switch {
	case strings.Contains(cfg.Endpoint, "4318"): // HTTP exporter
		exp, err = otlptracehttp.New(context.Background(),
			otlptracehttp.WithEndpoint(cfg.Endpoint),
			otlptracehttp.WithInsecure(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create HTTP OTLP exporter: %w", err)
		}
	case strings.Contains(cfg.Endpoint, "4317"): // gRPC exporter
		exp, err = otlptracegrpc.New(context.Background(),
			otlptracegrpc.WithEndpoint(cfg.Endpoint),
			otlptracegrpc.WithInsecure(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create gRPC OTLP exporter: %w", err)
		}
	default:
		return nil, fmt.Errorf("%w: got %s", ErrUnsupportedEndpoint, cfg.Endpoint)
	}

	return exp, nil
}

// createTraceProvider initializes a TracerProvider with the given exporter and configuration.
// The TracerProvider is responsible for managing tracers and exporting spans.
//
// Parameters:
// - cfg: Configuration containing service details and attributes.
// - exp: The span exporter to use.
//
// Returns:
// - *sdktrace.TracerProvider: The initialized tracer provider.
func createTraceProvider(cfg Config, exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,                                 // OpenTelemetry schema URL
			semconv.ServiceNameKey.String(cfg.AppName),        // Service name
			semconv.ServiceVersionKey.Float64(cfg.AppVersion), // Service version
			// Deployment environment (e.g., development, staging, production)
			semconv.DeploymentEnvironmentKey.String(cfg.AppEnv),
			// semconv.ServiceInstanceIDKey.String(_cfg.InstanceID), // Instance ID (useful for scaling scenarios)

			// Network information
			semconv.NetHostNameKey.String(cfg.AppHost), // ServiceHost name
			semconv.NetHostPortKey.Int(cfg.AppPort),    // ServicePort number

			// Dependency information
			// attribute.String("database", "PostgreSQL"), // Database in use
			// attribute.String("queue", "RabbitMQ"), // Message queue in use
			// attribute.String("cache", "Redis"), // Cache system in use

			// Other identifiers
			// attribute.String("team", "backend"), // Team responsible for the service
			// attribute.String("repository_url", "https://github.com/your-repo/service-name"), // Code repository URL
			// attribute.String("language", "Go"),                                              // Programming language
		)),
	)

	otel.SetTracerProvider(provider)

	return provider
}
