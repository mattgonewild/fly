// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeSnapshot volume snapshot
//
// swagger:model VolumeSnapshot
type VolumeSnapshot struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// digest
	Digest string `json:"digest,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this volume snapshot
func (m *VolumeSnapshot) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume snapshot based on context it is used
func (m *VolumeSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeSnapshot) UnmarshalBinary(b []byte) error {
	var res VolumeSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
