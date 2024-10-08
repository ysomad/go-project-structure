// Code generated by go-swagger; DO NOT EDIT.

package logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLogLevelHandlerFunc turns a function with the right signature into a get log level handler
type GetLogLevelHandlerFunc func(GetLogLevelParams) GetLogLevelResponder

// Handle executing the request and returning a response
func (fn GetLogLevelHandlerFunc) Handle(params GetLogLevelParams) GetLogLevelResponder {
	return fn(params)
}

// GetLogLevelHandler interface for that can handle valid get log level params
type GetLogLevelHandler interface {
	Handle(GetLogLevelParams) GetLogLevelResponder
}

// NewGetLogLevel creates a new http.Handler for the get log level operation
func NewGetLogLevel(ctx *middleware.Context, handler GetLogLevelHandler) *GetLogLevel {
	return &GetLogLevel{Context: ctx, Handler: handler}
}

/*
	GetLogLevel swagger:route GET /log-level logging getLogLevel

Текущий уровень логирования
*/
type GetLogLevel struct {
	Context *middleware.Context
	Handler GetLogLevelHandler
}

func (o *GetLogLevel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetLogLevelParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
