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

// NewMachinesDeleteParams creates a new MachinesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachinesDeleteParams() *MachinesDeleteParams {
	return &MachinesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachinesDeleteParamsWithTimeout creates a new MachinesDeleteParams object
// with the ability to set a timeout on a request.
func NewMachinesDeleteParamsWithTimeout(timeout time.Duration) *MachinesDeleteParams {
	return &MachinesDeleteParams{
		timeout: timeout,
	}
}

// NewMachinesDeleteParamsWithContext creates a new MachinesDeleteParams object
// with the ability to set a context for a request.
func NewMachinesDeleteParamsWithContext(ctx context.Context) *MachinesDeleteParams {
	return &MachinesDeleteParams{
		Context: ctx,
	}
}

// NewMachinesDeleteParamsWithHTTPClient creates a new MachinesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachinesDeleteParamsWithHTTPClient(client *http.Client) *MachinesDeleteParams {
	return &MachinesDeleteParams{
		HTTPClient: client,
	}
}

/*
MachinesDeleteParams contains all the parameters to send to the API endpoint

	for the machines delete operation.

	Typically these are written to a http.Request.
*/
type MachinesDeleteParams struct {

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

// WithDefaults hydrates default values in the machines delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesDeleteParams) WithDefaults() *MachinesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machines delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machines delete params
func (o *MachinesDeleteParams) WithTimeout(timeout time.Duration) *MachinesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machines delete params
func (o *MachinesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machines delete params
func (o *MachinesDeleteParams) WithContext(ctx context.Context) *MachinesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machines delete params
func (o *MachinesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machines delete params
func (o *MachinesDeleteParams) WithHTTPClient(client *http.Client) *MachinesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machines delete params
func (o *MachinesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the machines delete params
func (o *MachinesDeleteParams) WithAppName(appName string) *MachinesDeleteParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the machines delete params
func (o *MachinesDeleteParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithMachineID adds the machineID to the machines delete params
func (o *MachinesDeleteParams) WithMachineID(machineID string) *MachinesDeleteParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the machines delete params
func (o *MachinesDeleteParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *MachinesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
