// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// APIKey Api key
//
// swagger:model ApiKey
type APIKey string

// Validate validates this Api key
func (m APIKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api key based on context it is used
func (m APIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
