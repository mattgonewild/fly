// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIHTTPOptions api HTTP options
//
// swagger:model api.HTTPOptions
type APIHTTPOptions struct {

	// compress
	Compress bool `json:"compress,omitempty"`

	// h2 backend
	H2Backend bool `json:"h2_backend,omitempty"`

	// response
	Response *APIHTTPResponseOptions `json:"response,omitempty"`
}

// Validate validates this api HTTP options
func (m *APIHTTPOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIHTTPOptions) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api HTTP options based on the context it is used
func (m *APIHTTPOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIHTTPOptions) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {

		if swag.IsZero(m.Response) { // not required
			return nil
		}

		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIHTTPOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIHTTPOptions) UnmarshalBinary(b []byte) error {
	var res APIHTTPOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
