// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProductHandlerFunc turns a function with the right signature into a get product handler
type GetProductHandlerFunc func(GetProductParams) GetProductResponder

// Handle executing the request and returning a response
func (fn GetProductHandlerFunc) Handle(params GetProductParams) GetProductResponder {
	return fn(params)
}

// GetProductHandler interface for that can handle valid get product params
type GetProductHandler interface {
	Handle(GetProductParams) GetProductResponder
}

// NewGetProduct creates a new http.Handler for the get product operation
func NewGetProduct(ctx *middleware.Context, handler GetProductHandler) *GetProduct {
	return &GetProduct{Context: ctx, Handler: handler}
}

/*
	GetProduct swagger:route GET /v1/products/{product_id} product getProduct

Карточка товара
*/
type GetProduct struct {
	Context *middleware.Context
	Handler GetProductHandler
}

func (o *GetProduct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProductParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
