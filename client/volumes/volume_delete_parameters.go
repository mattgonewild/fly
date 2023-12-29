// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewVolumeDeleteParams creates a new VolumeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeDeleteParams() *VolumeDeleteParams {
	return &VolumeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeDeleteParamsWithTimeout creates a new VolumeDeleteParams object
// with the ability to set a timeout on a request.
func NewVolumeDeleteParamsWithTimeout(timeout time.Duration) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		timeout: timeout,
	}
}

// NewVolumeDeleteParamsWithContext creates a new VolumeDeleteParams object
// with the ability to set a context for a request.
func NewVolumeDeleteParamsWithContext(ctx context.Context) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		Context: ctx,
	}
}

// NewVolumeDeleteParamsWithHTTPClient creates a new VolumeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeDeleteParamsWithHTTPClient(client *http.Client) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		HTTPClient: client,
	}
}

/*
VolumeDeleteParams contains all the parameters to send to the API endpoint

	for the volume delete operation.

	Typically these are written to a http.Request.
*/
type VolumeDeleteParams struct {

	/* AppName.

	   Fly App Name
	*/
	AppName string

	/* VolumeID.

	   Volume ID
	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteParams) WithDefaults() *VolumeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume delete params
func (o *VolumeDeleteParams) WithTimeout(timeout time.Duration) *VolumeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume delete params
func (o *VolumeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume delete params
func (o *VolumeDeleteParams) WithContext(ctx context.Context) *VolumeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume delete params
func (o *VolumeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume delete params
func (o *VolumeDeleteParams) WithHTTPClient(client *http.Client) *VolumeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume delete params
func (o *VolumeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the volume delete params
func (o *VolumeDeleteParams) WithAppName(appName string) *VolumeDeleteParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the volume delete params
func (o *VolumeDeleteParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithVolumeID adds the volumeID to the volume delete params
func (o *VolumeDeleteParams) WithVolumeID(volumeID string) *VolumeDeleteParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the volume delete params
func (o *VolumeDeleteParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
