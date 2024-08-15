package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/ilyakaznacheev/cleanenv"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"

	"github.com/ysomad/go-project-structure/internal/config"
	"github.com/ysomad/go-project-structure/internal/gen/server/model"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/health"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/logging"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/product"
	serverhttp "github.com/ysomad/go-project-structure/internal/handler/http"
	"github.com/ysomad/go-project-structure/internal/handler/http/httpserver"
	v1 "github.com/ysomad/go-project-structure/internal/handler/http/v1"
	"github.com/ysomad/go-project-structure/internal/service"
	"github.com/ysomad/go-project-structure/internal/slogx"
	"github.com/ysomad/go-project-structure/internal/storage/postgres"
	"github.com/ysomad/go-project-structure/internal/storage/postgres/pgclient"
)

func Run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	// parse args
	var (
		confpath string
		conf     config.Server
		migrate  bool
	)

	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.StringVar(&confpath, "config", "./configs/server/local.toml", "config path")
	flags.StringVar(&conf.Metadata.Version, "version", "local-0", "service version")
	flags.StringVar(&conf.Metadata.InstanceID, "instance", "0", "service instance id")
	flags.Int64Var(&conf.Metadata.BuildTimestamp, "buildtime", time.Now().Unix(), "build timestamp UTC")
	flags.BoolVar(&migrate, "migrate", false, "run up migrations")

	if err := flags.Parse(args); err != nil {
		return fmt.Errorf("flags parse: %w", err)
	}

	// parse config
	if err := cleanenv.ReadConfig(confpath, &conf); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	// opentelemetry
	otelShutdown, err := setupOTelSDK(ctx, conf.Metadata)
	if err != nil {
		return err
	}
	defer func() {
		if err := otelShutdown(ctx); err != nil {
			slog.Warn("otel shutdown: " + err.Error())
		}
	}()

	// logger
	slogx.LevelVar.Set(slogx.ParseLevel(conf.Log.Level))

	otelHandler := otelslog.NewHandler("")
	logHandler := slogx.NewLevelFilter(otelHandler)
	slog.SetDefault(slog.New(logHandler))

	slog.Debug("starting with config", "config", conf)

	// migrations
	if migrate {
		err := runMigrations(ctx)
		if err != nil {
			return fmt.Errorf("goose: %w", err)
		}
	}

	// db
	pgcli, err := pgclient.New(ctx, conf.Postgres.URL)
	if err != nil {
		return fmt.Errorf("pgclient: %w", err)
	}

	// api
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return fmt.Errorf("swagger embed: %w", err)
	}

	api := operations.NewServerAPI(swaggerSpec)
	api.UseSwaggerUI()

	api.APIKeyAuth = func(key string) (any, error) {
		if key != conf.Log.APIKey {
			return nil, errors.New(401, "unauthenticated for invalid credentials")
		}
		return model.APIKey(key), nil
	}

	// domain
	productDB := postgres.NewProductStorage(pgcli)
	productSvc := service.NewProduct(productDB)

	// handler implementations
	api.HealthPingHandler = health.PingHandlerFunc(serverhttp.Ping)
	api.LoggingGetLogLevelHandler = logging.GetLogLevelHandlerFunc(serverhttp.GetLogLevel)
	api.LoggingUpdateLogLevelHandler = logging.UpdateLogLevelHandlerFunc(serverhttp.UpdateLogLevel)

	tracer := otel.GetTracerProvider().Tracer("")
	meter := otel.GetMeterProvider().Meter("")

	productHandlers := v1.NewProductHandlers(tracer, meter, productSvc)
	api.ProductListProductsV1Handler = product.ListProductsV1HandlerFunc(productHandlers.List)

	// provide opentelemetry tracing middleware to create spans on incoming requests
	swaggerHandler := api.Serve(withTracing)

	mux := http.NewServeMux()
	mux.Handle("/", swaggerHandler)

	if conf.Debug {
		attachPprof(mux)
	}

	// http
	srv := httpserver.New(ctx, mux, httpserver.WithPort(conf.Port))

	select {
	case err := <-srv.Notify():
		slogx.FatalContext(ctx, "httpserver: "+err.Error())
	case <-ctx.Done():
		slog.InfoContext(ctx, ctx.Err().Error())
	}

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
