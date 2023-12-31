// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIMachineInit api machine init
//
// swagger:model api.MachineInit
type APIMachineInit struct {

	// cmd
	Cmd []string `json:"cmd"`

	// entrypoint
	Entrypoint []string `json:"entrypoint"`

	// exec
	Exec []string `json:"exec"`

	// kernel args
	KernelArgs []string `json:"kernel_args"`

	// swap size mb
	SwapSizeMb int64 `json:"swap_size_mb,omitempty"`

	// tty
	Tty bool `json:"tty,omitempty"`
}

// Validate validates this api machine init
func (m *APIMachineInit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api machine init based on context it is used
func (m *APIMachineInit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIMachineInit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIMachineInit) UnmarshalBinary(b []byte) error {
	var res APIMachineInit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
