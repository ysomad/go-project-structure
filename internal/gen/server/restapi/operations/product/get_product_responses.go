// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/go-project-structure/internal/gen/server/model"
)

// GetProductOKCode is the HTTP code returned for type GetProductOK
const GetProductOKCode int = 200

/*
GetProductOK OK

swagger:response getProductOK
*/
type GetProductOK struct {

	/*
	  In: Body
	*/
	Payload *model.ProductCard `json:"body,omitempty"`
}

// NewGetProductOK creates GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {

	return &GetProductOK{}
}

// WithPayload adds the payload to the get product o k response
func (o *GetProductOK) WithPayload(payload *model.ProductCard) *GetProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product o k response
func (o *GetProductOK) SetPayload(payload *model.ProductCard) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetProductOK) GetProductResponder() {}

// GetProductBadRequestCode is the HTTP code returned for type GetProductBadRequest
const GetProductBadRequestCode int = 400

/*
GetProductBadRequest Bad Request

swagger:response getProductBadRequest
*/
type GetProductBadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewGetProductBadRequest creates GetProductBadRequest with default headers values
func NewGetProductBadRequest() *GetProductBadRequest {

	return &GetProductBadRequest{}
}

// WithPayload adds the payload to the get product bad request response
func (o *GetProductBadRequest) WithPayload(payload *model.Error) *GetProductBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product bad request response
func (o *GetProductBadRequest) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetProductBadRequest) GetProductResponder() {}

type GetProductNotImplementedResponder struct {
	middleware.Responder
}

func (*GetProductNotImplementedResponder) GetProductResponder() {}

func GetProductNotImplemented() GetProductResponder {
	return &GetProductNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetProduct has not yet been implemented",
		),
	}
}

type GetProductResponder interface {
	middleware.Responder
	GetProductResponder()
}
