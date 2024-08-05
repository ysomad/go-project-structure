// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListProductsResponse list products response
//
// swagger:model ListProductsResponse
type ListProductsResponse struct {

	// next page token
	NextPageToken string `json:"next_page_token,omitempty"`

	// products
	Products []*Product `json:"products"`

	// total pages
	TotalPages int32 `json:"total_pages,omitempty"`
}

// Validate validates this list products response
func (m *ListProductsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListProductsResponse) validateProducts(formats strfmt.Registry) error {
	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list products response based on the context it is used
func (m *ListProductsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProducts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListProductsResponse) contextValidateProducts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Products); i++ {

		if m.Products[i] != nil {

			if swag.IsZero(m.Products[i]) { // not required
				return nil
			}

			if err := m.Products[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListProductsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListProductsResponse) UnmarshalBinary(b []byte) error {
	var res ListProductsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
