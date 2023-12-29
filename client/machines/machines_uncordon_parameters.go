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

// NewMachinesUncordonParams creates a new MachinesUncordonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachinesUncordonParams() *MachinesUncordonParams {
	return &MachinesUncordonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachinesUncordonParamsWithTimeout creates a new MachinesUncordonParams object
// with the ability to set a timeout on a request.
func NewMachinesUncordonParamsWithTimeout(timeout time.Duration) *MachinesUncordonParams {
	return &MachinesUncordonParams{
		timeout: timeout,
	}
}

// NewMachinesUncordonParamsWithContext creates a new MachinesUncordonParams object
// with the ability to set a context for a request.
func NewMachinesUncordonParamsWithContext(ctx context.Context) *MachinesUncordonParams {
	return &MachinesUncordonParams{
		Context: ctx,
	}
}

// NewMachinesUncordonParamsWithHTTPClient creates a new MachinesUncordonParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachinesUncordonParamsWithHTTPClient(client *http.Client) *MachinesUncordonParams {
	return &MachinesUncordonParams{
		HTTPClient: client,
	}
}

/*
MachinesUncordonParams contains all the parameters to send to the API endpoint

	for the machines uncordon operation.

	Typically these are written to a http.Request.
*/
type MachinesUncordonParams struct {

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

// WithDefaults hydrates default values in the machines uncordon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesUncordonParams) WithDefaults() *MachinesUncordonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machines uncordon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesUncordonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machines uncordon params
func (o *MachinesUncordonParams) WithTimeout(timeout time.Duration) *MachinesUncordonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machines uncordon params
func (o *MachinesUncordonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machines uncordon params
func (o *MachinesUncordonParams) WithContext(ctx context.Context) *MachinesUncordonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machines uncordon params
func (o *MachinesUncordonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machines uncordon params
func (o *MachinesUncordonParams) WithHTTPClient(client *http.Client) *MachinesUncordonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machines uncordon params
func (o *MachinesUncordonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the machines uncordon params
func (o *MachinesUncordonParams) WithAppName(appName string) *MachinesUncordonParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the machines uncordon params
func (o *MachinesUncordonParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithMachineID adds the machineID to the machines uncordon params
func (o *MachinesUncordonParams) WithMachineID(machineID string) *MachinesUncordonParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the machines uncordon params
func (o *MachinesUncordonParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *MachinesUncordonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
