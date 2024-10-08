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

// ListProductsV1OKCode is the HTTP code returned for type ListProductsV1OK
const ListProductsV1OKCode int = 200

/*
ListProductsV1OK OK

swagger:response listProductsV1OK
*/
type ListProductsV1OK struct {

	/*
	  In: Body
	*/
	Payload *model.ListProductsResponse `json:"body,omitempty"`
}

// NewListProductsV1OK creates ListProductsV1OK with default headers values
func NewListProductsV1OK() *ListProductsV1OK {

	return &ListProductsV1OK{}
}

// WithPayload adds the payload to the list products v1 o k response
func (o *ListProductsV1OK) WithPayload(payload *model.ListProductsResponse) *ListProductsV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list products v1 o k response
func (o *ListProductsV1OK) SetPayload(payload *model.ListProductsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProductsV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListProductsV1OK) ListProductsV1Responder() {}

// ListProductsV1BadRequestCode is the HTTP code returned for type ListProductsV1BadRequest
const ListProductsV1BadRequestCode int = 400

/*
ListProductsV1BadRequest Bad Request

swagger:response listProductsV1BadRequest
*/
type ListProductsV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewListProductsV1BadRequest creates ListProductsV1BadRequest with default headers values
func NewListProductsV1BadRequest() *ListProductsV1BadRequest {

	return &ListProductsV1BadRequest{}
}

// WithPayload adds the payload to the list products v1 bad request response
func (o *ListProductsV1BadRequest) WithPayload(payload *model.Error) *ListProductsV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list products v1 bad request response
func (o *ListProductsV1BadRequest) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProductsV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListProductsV1BadRequest) ListProductsV1Responder() {}

type ListProductsV1NotImplementedResponder struct {
	middleware.Responder
}

func (*ListProductsV1NotImplementedResponder) ListProductsV1Responder() {}

func ListProductsV1NotImplemented() ListProductsV1Responder {
	return &ListProductsV1NotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.ListProductsV1 has not yet been implemented",
		),
	}
}

type ListProductsV1Responder interface {
	middleware.Responder
	ListProductsV1Responder()
}
