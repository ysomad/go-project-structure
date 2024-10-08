// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/go-project-structure/internal/gen/server/model"
)

// PingOKCode is the HTTP code returned for type PingOK
const PingOKCode int = 200

/*
PingOK OK

swagger:response pingOK
*/
type PingOK struct {

	/*
	  In: Body
	*/
	Payload *model.PingResponse `json:"body,omitempty"`
}

// NewPingOK creates PingOK with default headers values
func NewPingOK() *PingOK {

	return &PingOK{}
}

// WithPayload adds the payload to the ping o k response
func (o *PingOK) WithPayload(payload *model.PingResponse) *PingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ping o k response
func (o *PingOK) SetPayload(payload *model.PingResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PingOK) PingResponder() {}

type PingNotImplementedResponder struct {
	middleware.Responder
}

func (*PingNotImplementedResponder) PingResponder() {}

func PingNotImplemented() PingResponder {
	return &PingNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Ping has not yet been implemented",
		),
	}
}

type PingResponder interface {
	middleware.Responder
	PingResponder()
}
