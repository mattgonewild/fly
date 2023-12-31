// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Lease lease
//
// swagger:model Lease
type Lease struct {

	// Description or reason for the Lease.
	Description string `json:"description,omitempty"`

	// ExpiresAt is the unix timestamp in UTC to denote when the Lease will no longer be valid.
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// Nonce is the unique ID autogenerated and associated with the Lease.
	Nonce string `json:"nonce,omitempty"`

	// Owner is the user identifier which acquired the Lease.
	Owner string `json:"owner,omitempty"`

	// Machine version
	Version string `json:"version,omitempty"`
}

// Validate validates this lease
func (m *Lease) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lease based on context it is used
func (m *Lease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Lease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Lease) UnmarshalBinary(b []byte) error {
	var res Lease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
