// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OfferMerchant offer merchant
//
// swagger:model OfferMerchant
type OfferMerchant struct {

	// icon url
	// Format: uri
	IconURL strfmt.URI `json:"icon_url,omitempty"`

	// logo text
	LogoText string `json:"logo_text,omitempty"`

	// logo url
	// Format: uri
	LogoURL strfmt.URI `json:"logo_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// official
	Official bool `json:"official,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this offer merchant
func (m *OfferMerchant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferMerchant) validateIconURL(formats strfmt.Registry) error {
	if swag.IsZero(m.IconURL) { // not required
		return nil
	}

	if err := validate.FormatOf("icon_url", "body", "uri", m.IconURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OfferMerchant) validateLogoURL(formats strfmt.Registry) error {
	if swag.IsZero(m.LogoURL) { // not required
		return nil
	}

	if err := validate.FormatOf("logo_url", "body", "uri", m.LogoURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this offer merchant based on context it is used
func (m *OfferMerchant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OfferMerchant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferMerchant) UnmarshalBinary(b []byte) error {
	var res OfferMerchant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
