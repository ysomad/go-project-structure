// Code generated by go-swagger; DO NOT EDIT.

package logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// UpdateLogLevelNoContentCode is the HTTP code returned for type UpdateLogLevelNoContent
const UpdateLogLevelNoContentCode int = 204

/*
UpdateLogLevelNoContent No Content

swagger:response updateLogLevelNoContent
*/
type UpdateLogLevelNoContent struct {
}

// NewUpdateLogLevelNoContent creates UpdateLogLevelNoContent with default headers values
func NewUpdateLogLevelNoContent() *UpdateLogLevelNoContent {

	return &UpdateLogLevelNoContent{}
}

// WriteResponse to the client
func (o *UpdateLogLevelNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateLogLevelNoContent) UpdateLogLevelResponder() {}

type UpdateLogLevelNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateLogLevelNotImplementedResponder) UpdateLogLevelResponder() {}

func UpdateLogLevelNotImplemented() UpdateLogLevelResponder {
	return &UpdateLogLevelNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateLogLevel has not yet been implemented",
		),
	}
}

type UpdateLogLevelResponder interface {
	middleware.Responder
	UpdateLogLevelResponder()
}
