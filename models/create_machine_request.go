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

// CreateMachineRequest create machine request
//
// swagger:model CreateMachineRequest
type CreateMachineRequest struct {

	// An object defining the Machine configuration
	Config struct {
		APIMachineConfig
	} `json:"config,omitempty"`

	// lease ttl
	LeaseTTL int64 `json:"lease_ttl,omitempty"`

	// lsvd
	Lsvd bool `json:"lsvd,omitempty"`

	// Unique name for this Machine. If omitted, one is generated for you
	Name string `json:"name,omitempty"`

	// The target region. Omitting this param launches in the same region as your WireGuard peer connection (somewhere near you).
	Region string `json:"region,omitempty"`

	// skip launch
	SkipLaunch bool `json:"skip_launch,omitempty"`

	// skip service registration
	SkipServiceRegistration bool `json:"skip_service_registration,omitempty"`
}

// Validate validates this create machine request
func (m *CreateMachineRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineRequest) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this create machine request based on the context it is used
func (m *CreateMachineRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineRequest) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMachineRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMachineRequest) UnmarshalBinary(b []byte) error {
	var res CreateMachineRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
