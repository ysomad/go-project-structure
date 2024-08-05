// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/health"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/logging"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/product"
)

//go:generate swagger generate server --target ../../server --name Server --spec ../../../../api/swagger2.yaml --model-package model --principal interface{} --exclude-main --strict-responders

func configureFlags(api *operations.ServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-Key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (apiKey) X-API-Key from header param [X-API-Key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.LoggingGetLogLevelHandler == nil {
		api.LoggingGetLogLevelHandler = logging.GetLogLevelHandlerFunc(func(params logging.GetLogLevelParams) logging.GetLogLevelResponder {
			return logging.GetLogLevelNotImplemented()
		})
	}
	if api.ProductGetProductHandler == nil {
		api.ProductGetProductHandler = product.GetProductHandlerFunc(func(params product.GetProductParams) product.GetProductResponder {
			return product.GetProductNotImplemented()
		})
	}
	if api.ProductListProductsV1Handler == nil {
		api.ProductListProductsV1Handler = product.ListProductsV1HandlerFunc(func(params product.ListProductsV1Params) product.ListProductsV1Responder {
			return product.ListProductsV1NotImplemented()
		})
	}
	if api.HealthPingHandler == nil {
		api.HealthPingHandler = health.PingHandlerFunc(func(params health.PingParams) health.PingResponder {
			return health.PingNotImplemented()
		})
	}
	if api.LoggingUpdateLogLevelHandler == nil {
		api.LoggingUpdateLogLevelHandler = logging.UpdateLogLevelHandlerFunc(func(params logging.UpdateLogLevelParams, principal interface{}) logging.UpdateLogLevelResponder {
			return logging.UpdateLogLevelNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
