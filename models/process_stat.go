// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessStat process stat
//
// swagger:model ProcessStat
type ProcessStat struct {

	// command
	Command string `json:"command,omitempty"`

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// directory
	Directory string `json:"directory,omitempty"`

	// listen sockets
	ListenSockets []*ListenSocket `json:"listen_sockets"`

	// pid
	Pid int64 `json:"pid,omitempty"`

	// rss
	Rss int64 `json:"rss,omitempty"`

	// rtime
	Rtime int64 `json:"rtime,omitempty"`

	// stime
	Stime int64 `json:"stime,omitempty"`
}

// Validate validates this process stat
func (m *ProcessStat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListenSockets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessStat) validateListenSockets(formats strfmt.Registry) error {
	if swag.IsZero(m.ListenSockets) { // not required
		return nil
	}

	for i := 0; i < len(m.ListenSockets); i++ {
		if swag.IsZero(m.ListenSockets[i]) { // not required
			continue
		}

		if m.ListenSockets[i] != nil {
			if err := m.ListenSockets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listen_sockets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listen_sockets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this process stat based on the context it is used
func (m *ProcessStat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListenSockets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessStat) contextValidateListenSockets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ListenSockets); i++ {

		if m.ListenSockets[i] != nil {

			if swag.IsZero(m.ListenSockets[i]) { // not required
				return nil
			}

			if err := m.ListenSockets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listen_sockets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listen_sockets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessStat) UnmarshalBinary(b []byte) error {
	var res ProcessStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
