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

// NewMachinesShowMetadataParams creates a new MachinesShowMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachinesShowMetadataParams() *MachinesShowMetadataParams {
	return &MachinesShowMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachinesShowMetadataParamsWithTimeout creates a new MachinesShowMetadataParams object
// with the ability to set a timeout on a request.
func NewMachinesShowMetadataParamsWithTimeout(timeout time.Duration) *MachinesShowMetadataParams {
	return &MachinesShowMetadataParams{
		timeout: timeout,
	}
}

// NewMachinesShowMetadataParamsWithContext creates a new MachinesShowMetadataParams object
// with the ability to set a context for a request.
func NewMachinesShowMetadataParamsWithContext(ctx context.Context) *MachinesShowMetadataParams {
	return &MachinesShowMetadataParams{
		Context: ctx,
	}
}

// NewMachinesShowMetadataParamsWithHTTPClient creates a new MachinesShowMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachinesShowMetadataParamsWithHTTPClient(client *http.Client) *MachinesShowMetadataParams {
	return &MachinesShowMetadataParams{
		HTTPClient: client,
	}
}

/*
MachinesShowMetadataParams contains all the parameters to send to the API endpoint

	for the machines show metadata operation.

	Typically these are written to a http.Request.
*/
type MachinesShowMetadataParams struct {

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

// WithDefaults hydrates default values in the machines show metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesShowMetadataParams) WithDefaults() *MachinesShowMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machines show metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinesShowMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machines show metadata params
func (o *MachinesShowMetadataParams) WithTimeout(timeout time.Duration) *MachinesShowMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machines show metadata params
func (o *MachinesShowMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machines show metadata params
func (o *MachinesShowMetadataParams) WithContext(ctx context.Context) *MachinesShowMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machines show metadata params
func (o *MachinesShowMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machines show metadata params
func (o *MachinesShowMetadataParams) WithHTTPClient(client *http.Client) *MachinesShowMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machines show metadata params
func (o *MachinesShowMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the machines show metadata params
func (o *MachinesShowMetadataParams) WithAppName(appName string) *MachinesShowMetadataParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the machines show metadata params
func (o *MachinesShowMetadataParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithMachineID adds the machineID to the machines show metadata params
func (o *MachinesShowMetadataParams) WithMachineID(machineID string) *MachinesShowMetadataParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the machines show metadata params
func (o *MachinesShowMetadataParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *MachinesShowMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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