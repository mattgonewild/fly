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

	"github.com/mattgonewild/fly/models"
)

// NewVolumesUpdateParams creates a new VolumesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumesUpdateParams() *VolumesUpdateParams {
	return &VolumesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumesUpdateParamsWithTimeout creates a new VolumesUpdateParams object
// with the ability to set a timeout on a request.
func NewVolumesUpdateParamsWithTimeout(timeout time.Duration) *VolumesUpdateParams {
	return &VolumesUpdateParams{
		timeout: timeout,
	}
}

// NewVolumesUpdateParamsWithContext creates a new VolumesUpdateParams object
// with the ability to set a context for a request.
func NewVolumesUpdateParamsWithContext(ctx context.Context) *VolumesUpdateParams {
	return &VolumesUpdateParams{
		Context: ctx,
	}
}

// NewVolumesUpdateParamsWithHTTPClient creates a new VolumesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumesUpdateParamsWithHTTPClient(client *http.Client) *VolumesUpdateParams {
	return &VolumesUpdateParams{
		HTTPClient: client,
	}
}

/*
VolumesUpdateParams contains all the parameters to send to the API endpoint

	for the volumes update operation.

	Typically these are written to a http.Request.
*/
type VolumesUpdateParams struct {

	/* AppName.

	   Fly App Name
	*/
	AppName string

	/* Request.

	   Request body
	*/
	Request *models.UpdateVolumeRequest

	/* VolumeID.

	   Volume ID
	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volumes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumesUpdateParams) WithDefaults() *VolumesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volumes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volumes update params
func (o *VolumesUpdateParams) WithTimeout(timeout time.Duration) *VolumesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volumes update params
func (o *VolumesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volumes update params
func (o *VolumesUpdateParams) WithContext(ctx context.Context) *VolumesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volumes update params
func (o *VolumesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volumes update params
func (o *VolumesUpdateParams) WithHTTPClient(client *http.Client) *VolumesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volumes update params
func (o *VolumesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the volumes update params
func (o *VolumesUpdateParams) WithAppName(appName string) *VolumesUpdateParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the volumes update params
func (o *VolumesUpdateParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithRequest adds the request to the volumes update params
func (o *VolumesUpdateParams) WithRequest(request *models.UpdateVolumeRequest) *VolumesUpdateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the volumes update params
func (o *VolumesUpdateParams) SetRequest(request *models.UpdateVolumeRequest) {
	o.Request = request
}

// WithVolumeID adds the volumeID to the volumes update params
func (o *VolumesUpdateParams) WithVolumeID(volumeID string) *VolumesUpdateParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the volumes update params
func (o *VolumesUpdateParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *VolumesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
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
