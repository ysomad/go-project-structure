// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PingResponse ping response
//
// swagger:model PingResponse
type PingResponse struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this ping response
func (m *PingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ping response based on context it is used
func (m *PingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingResponse) UnmarshalBinary(b []byte) error {
	var res PingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
