package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"

	"github.com/ysomad/go-project-structure/internal/config"
	"github.com/ysomad/go-project-structure/internal/gen/server/model"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/health"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/logging"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/product"
	"github.com/ysomad/go-project-structure/internal/handler/http"
	v1 "github.com/ysomad/go-project-structure/internal/handler/http/v1"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

func Run(conf config.Server) error {
	// opentelemetry
	otelShutdown, err := setupOTelSDK(context.TODO(), conf.Metadata)
	if err != nil {
		return err
	}
	defer func() {
		if err := otelShutdown(context.Background()); err != nil {
			slogx.Fatal("otel shutdown: %w", err)
		}
	}()

	// logger
	otelHandler := otelslog.NewHandler("")
	levelHandler := slogx.NewLevelFilter(otelHandler, conf.Log.Level)
	slog.SetDefault(slog.New(levelHandler))

	slog.Debug("starting with config", "config", conf)

	// global providers
	tracer := otel.GetTracerProvider().Tracer("")
	meter := otel.GetMeterProvider().Meter("")

	// api
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return fmt.Errorf("swagger embed: %w", err)
	}

	api := operations.NewServerAPI(swaggerSpec)
	api.UseSwaggerUI()

	api.APIKeyAuth = func(key string) (any, error) {
		if key == conf.Log.APIKey {
			return model.APIKey(key), nil
		}
		return nil, errors.New(401, "unauthenticated for invalid credentials")
	}

	// handler implementations
	api.HealthPingHandler = health.PingHandlerFunc(http.Ping)
	api.LoggingGetLogLevelHandler = logging.GetLogLevelHandlerFunc(http.GetLogLevel)
	api.LoggingUpdateLogLevelHandler = logging.UpdateLogLevelHandlerFunc(http.UpdateLogLevel)

	productHandlers := v1.NewProductHandlers(tracer, meter)
	api.ProductListProductsV1Handler = product.ListProductsV1HandlerFunc(productHandlers.List)

	server := restapi.NewServer(api)
	defer server.Shutdown() //nolint:errcheck // never fails

	server.Port = conf.Port

	// provide opentelemetry tracing middleware to create spans on incoming requests
	handler := api.Serve(withTracing)

	server.SetHandler(handler)

	if err := server.Serve(); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}
