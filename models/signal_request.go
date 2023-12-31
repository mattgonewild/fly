// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignalRequest signal request
//
// swagger:model SignalRequest
type SignalRequest struct {

	// signal
	// Enum: [SIGABRT SIGALRM SIGFPE SIGHUP SIGILL SIGINT SIGKILL SIGPIPE SIGQUIT SIGSEGV SIGTERM SIGTRAP SIGUSR1]
	Signal string `json:"signal,omitempty"`
}

// Validate validates this signal request
func (m *SignalRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var signalRequestTypeSignalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SIGABRT","SIGALRM","SIGFPE","SIGHUP","SIGILL","SIGINT","SIGKILL","SIGPIPE","SIGQUIT","SIGSEGV","SIGTERM","SIGTRAP","SIGUSR1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signalRequestTypeSignalPropEnum = append(signalRequestTypeSignalPropEnum, v)
	}
}

const (

	// SignalRequestSignalSIGABRT captures enum value "SIGABRT"
	SignalRequestSignalSIGABRT string = "SIGABRT"

	// SignalRequestSignalSIGALRM captures enum value "SIGALRM"
	SignalRequestSignalSIGALRM string = "SIGALRM"

	// SignalRequestSignalSIGFPE captures enum value "SIGFPE"
	SignalRequestSignalSIGFPE string = "SIGFPE"

	// SignalRequestSignalSIGHUP captures enum value "SIGHUP"
	SignalRequestSignalSIGHUP string = "SIGHUP"

	// SignalRequestSignalSIGILL captures enum value "SIGILL"
	SignalRequestSignalSIGILL string = "SIGILL"

	// SignalRequestSignalSIGINT captures enum value "SIGINT"
	SignalRequestSignalSIGINT string = "SIGINT"

	// SignalRequestSignalSIGKILL captures enum value "SIGKILL"
	SignalRequestSignalSIGKILL string = "SIGKILL"

	// SignalRequestSignalSIGPIPE captures enum value "SIGPIPE"
	SignalRequestSignalSIGPIPE string = "SIGPIPE"

	// SignalRequestSignalSIGQUIT captures enum value "SIGQUIT"
	SignalRequestSignalSIGQUIT string = "SIGQUIT"

	// SignalRequestSignalSIGSEGV captures enum value "SIGSEGV"
	SignalRequestSignalSIGSEGV string = "SIGSEGV"

	// SignalRequestSignalSIGTERM captures enum value "SIGTERM"
	SignalRequestSignalSIGTERM string = "SIGTERM"

	// SignalRequestSignalSIGTRAP captures enum value "SIGTRAP"
	SignalRequestSignalSIGTRAP string = "SIGTRAP"

	// SignalRequestSignalSIGUSR1 captures enum value "SIGUSR1"
	SignalRequestSignalSIGUSR1 string = "SIGUSR1"
)

// prop value enum
func (m *SignalRequest) validateSignalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, signalRequestTypeSignalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SignalRequest) validateSignal(formats strfmt.Registry) error {
	if swag.IsZero(m.Signal) { // not required
		return nil
	}

	// value enum
	if err := m.validateSignalEnum("signal", "body", m.Signal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signal request based on context it is used
func (m *SignalRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignalRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignalRequest) UnmarshalBinary(b []byte) error {
	var res SignalRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
