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

// NewCreateVolumeSnapshotParams creates a new CreateVolumeSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVolumeSnapshotParams() *CreateVolumeSnapshotParams {
	return &CreateVolumeSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVolumeSnapshotParamsWithTimeout creates a new CreateVolumeSnapshotParams object
// with the ability to set a timeout on a request.
func NewCreateVolumeSnapshotParamsWithTimeout(timeout time.Duration) *CreateVolumeSnapshotParams {
	return &CreateVolumeSnapshotParams{
		timeout: timeout,
	}
}

// NewCreateVolumeSnapshotParamsWithContext creates a new CreateVolumeSnapshotParams object
// with the ability to set a context for a request.
func NewCreateVolumeSnapshotParamsWithContext(ctx context.Context) *CreateVolumeSnapshotParams {
	return &CreateVolumeSnapshotParams{
		Context: ctx,
	}
}

// NewCreateVolumeSnapshotParamsWithHTTPClient creates a new CreateVolumeSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVolumeSnapshotParamsWithHTTPClient(client *http.Client) *CreateVolumeSnapshotParams {
	return &CreateVolumeSnapshotParams{
		HTTPClient: client,
	}
}

/*
CreateVolumeSnapshotParams contains all the parameters to send to the API endpoint

	for the create volume snapshot operation.

	Typically these are written to a http.Request.
*/
type CreateVolumeSnapshotParams struct {

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

// WithDefaults hydrates default values in the create volume snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVolumeSnapshotParams) WithDefaults() *CreateVolumeSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create volume snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVolumeSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) WithTimeout(timeout time.Duration) *CreateVolumeSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) WithContext(ctx context.Context) *CreateVolumeSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) WithHTTPClient(client *http.Client) *CreateVolumeSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) WithAppName(appName string) *CreateVolumeSnapshotParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithVolumeID adds the volumeID to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) WithVolumeID(volumeID string) *CreateVolumeSnapshotParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the create volume snapshot params
func (o *CreateVolumeSnapshotParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVolumeSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
