package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ysomad/go-project-structure/internal/config"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context, meta config.Metadata) (func(context.Context) error, error) {
	var (
		shutdownFuncs []func(context.Context) error
		err           error
	)

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown := func(ctx context.Context) error {
		slog.DebugContext(ctx, "otel cleanup")
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	res := newResource(meta.ServiceName, meta.Version, meta.InstanceID, meta.BuildTimestamp)

	// tracing
	traceProvider, err := newTracerProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)
	otel.SetTracerProvider(traceProvider)

	// metrics
	meterProvider, err := newMeterProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// logging
	loggerProvider, err := newLoggerProvider(res)
	if err != nil {
		handleErr(err)
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return shutdown, nil
}

// withTracing wraps each handler in otelhttp handler to create span on incoming requests.
func withTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		otelhttp.NewHandler(next, middleware.MatchedRouteFrom(r).PathPattern).ServeHTTP(w, r)
	})
}

func newPropagator() propagation.TextMapPropagator { //nolint:ireturn // opentelemetry spec
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

// https://opentelemetry.io/docs/languages/go/resources/
func newResource(serviceName, version, instanceID string, buildTime int64) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(version),
		semconv.ServiceInstanceIDKey.String(instanceID),
		attribute.Int64("service.build_time", buildTime),
	)
}

func newTracerProvider(ctx context.Context, res *resource.Resource) (*trace.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("tracer provider: %w", err)
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res),
	), nil
}

func newMeterProvider(ctx context.Context, res *resource.Resource) (*metric.MeterProvider, error) {
	exp, err := otlpmetrichttp.New(ctx, otlpmetrichttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("meter provider: %w", err)
	}
	reader := metric.NewPeriodicReader(exp)
	return metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(reader),
	), nil
}

func newLoggerProvider(res *resource.Resource) (*log.LoggerProvider, error) {
	exp, err := stdoutlog.New()
	if err != nil {
		return nil, fmt.Errorf("log provider: %w", err)
	}
	processor := log.NewBatchProcessor(exp)
	return log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(processor),
	), nil
}
