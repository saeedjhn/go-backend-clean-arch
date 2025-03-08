package otelcollector

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OpenTelemetry encapsulates the configuration and methods
// to interact with OpenTelemetry metrics using a MeterProvider and counters.
type OpenTelemetry struct {
	options          Options
	bucketBoundaries []float64
	meter            metric.Meter
	meterProvider    *sdkmetric.MeterProvider
	counterCache     sync.Map // Cache for storing counters by name to avoid recreation.
}

// New creates a new instance of OpenTelemetry with the provided context and configuration.
func New(ops Options) *OpenTelemetry {
	return &OpenTelemetry{
		options:       ops,
		meter:         nil,
		meterProvider: nil,
		counterCache:  sync.Map{},
	}
}

// Configure sets up the OpenTelemetry instance by initializing the OTLP exporter,
// establishing a connection, and configuring the MeterProvider.
func (o *OpenTelemetry) Configure() error {
	if len(o.options.Config.Endpoint) == 0 {
		return ErrEndpointRequired
	}

	conn, err := createConn(o.options.Config.Endpoint)
	if err != nil {
		return err
	}

	exp, err := createExporter(conn, o.options.Config.Timeout)
	if err != nil {
		return err
	}

	mp := createMeterProvider(exp, o.options.AppInfo)

	o.meter = mp.Meter(o.options.AppInfo.Name)
	o.meterProvider = mp

	return nil
}

// Shutdown gracefully shuts down the MeterProvider and releases resources.
func (o *OpenTelemetry) Shutdown(ctx context.Context) error {
	if o.meter == nil {
		return nil // No tracer initialized, no need to shut down
	}

	if o.meterProvider == nil {
		return nil // No meter provider initialized, nothing to shut down
	}

	if err := o.meterProvider.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down metric provider: %w", err)
	}

	return nil
}

func (o *OpenTelemetry) WithBucketBoundaries(bounds []float64) contract.Collector {
	o.bucketBoundaries = bounds

	return o
}

func (o *OpenTelemetry) setAttrs(attrs []map[string]interface{}) []attribute.KeyValue {
	var otelAttrs []attribute.KeyValue

	// Convert attributes into OpenTelemetry format.
	if len(attrs) != 0 {
		for _, attr := range attrs {
			for k, v := range attr {
				otelAttrs = append(otelAttrs, attribute.String(k, fmt.Sprintf("%#v", v)))
			}
		}
	}

	return otelAttrs
}

// createConn establishes a gRPC connection to the specified OTLP endpoint.
func createConn(otlpEndpoint string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(otlpEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, nil
}

// createExporter initializes a gRPC-based OTLP metric exporter.
func createExporter(conn *grpc.ClientConn, timeout time.Duration) (*otlpmetricgrpc.Exporter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	exp, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithGRPCConn(conn),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC OTLP exporter: %w", err)
	}

	return exp, nil
}

// createMeterProvider sets up a MeterProvider for managing metric data.
func createMeterProvider(exp *otlpmetricgrpc.Exporter, appInfo AppInfo) *sdkmetric.MeterProvider {
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)),
		sdkmetric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL, // OpenTelemetry schema URL
			semconv.ServiceNameKey.String(appInfo.Name),       //  name
			semconv.ServiceVersionKey.String(appInfo.Version), //  version
			// Deployment environment (e.g., development, staging, production)
			semconv.DeploymentEnvironmentKey.String(appInfo.Env),
			// semconv.ServiceInstanceIDKey.String(appInfo.InstanceID), // Instance ID (useful for scaling scenarios)

			// Network information
			semconv.NetHostNameKey.String(appInfo.Host), // Host name
			semconv.NetHostPortKey.String(appInfo.Port), // Port number
		)),
	)

	otel.SetMeterProvider(provider)

	return provider
}
