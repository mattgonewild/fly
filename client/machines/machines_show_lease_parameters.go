// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewMachinesShowLeaseParams creates a new MachinesShowLeaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachinesShowLeaseParams() *MachinesShowLeaseParams {
	return &MachinesShowLeaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachinesShowLeaseParamsWithTimeout creates a new MachinesShowLeaseParams object
// with the ability to set a timeout on a request.
func NewMachinesShowLeaseParamsWithTimeout(timeout time.Duration) *MachinesShowLeaseParams {
	return &MachinesShowLeaseParams{
		timeout: timeout,
	}
}

// NewMachinesShowLeaseParamsWithContext creates a new MachinesShowLeaseParams object
// with the ability to set a context for a request.
func NewMachinesShowLeaseParamsWithContext(ctx context.Context) *MachinesShowLeaseParams {
	return &MachinesShowLeaseParams{
		Context: ctx,
	}
}

// NewMachinesShowLeaseParamsWithHTTPClient creates a new MachinesShowLeaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachinesShowLeaseParamsWithHTTPClient(client *http.Client) *MachinesShowLeaseParams {
	return &MachinesShowLeaseParams{
		HTTPClient: client,
	}
}

/*
MachinesShowLeaseParams contains all the parameters to send to the API endpoint

	for the machines show lease operation.

	Typically these are written to a http.Request.
*/
type MachinesShowLeaseParams struct {

	/* AppName.

	   Fly App Name
	*/
	AppName string

	/* MachineID.

	   Machine ID
	*/
	MachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the machines show lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesShowLeaseParams) WithDefaults() *MachinesShowLeaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machines show lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesShowLeaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machines show lease params
func (o *MachinesShowLeaseParams) WithTimeout(timeout time.Duration) *MachinesShowLeaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machines show lease params
func (o *MachinesShowLeaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machines show lease params
func (o *MachinesShowLeaseParams) WithContext(ctx context.Context) *MachinesShowLeaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machines show lease params
func (o *MachinesShowLeaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machines show lease params
func (o *MachinesShowLeaseParams) WithHTTPClient(client *http.Client) *MachinesShowLeaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machines show lease params
func (o *MachinesShowLeaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the machines show lease params
func (o *MachinesShowLeaseParams) WithAppName(appName string) *MachinesShowLeaseParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the machines show lease params
func (o *MachinesShowLeaseParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithMachineID adds the machineID to the machines show lease params
func (o *MachinesShowLeaseParams) WithMachineID(machineID string) *MachinesShowLeaseParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the machines show lease params
func (o *MachinesShowLeaseParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *MachinesShowLeaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	// path param machine_id
	if err := r.SetPathParam("machine_id", o.MachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
