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

// NewVolumesCreateParams creates a new VolumesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumesCreateParams() *VolumesCreateParams {
	return &VolumesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumesCreateParamsWithTimeout creates a new VolumesCreateParams object
// with the ability to set a timeout on a request.
func NewVolumesCreateParamsWithTimeout(timeout time.Duration) *VolumesCreateParams {
	return &VolumesCreateParams{
		timeout: timeout,
	}
}

// NewVolumesCreateParamsWithContext creates a new VolumesCreateParams object
// with the ability to set a context for a request.
func NewVolumesCreateParamsWithContext(ctx context.Context) *VolumesCreateParams {
	return &VolumesCreateParams{
		Context: ctx,
	}
}

// NewVolumesCreateParamsWithHTTPClient creates a new VolumesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumesCreateParamsWithHTTPClient(client *http.Client) *VolumesCreateParams {
	return &VolumesCreateParams{
		HTTPClient: client,
	}
}

/*
VolumesCreateParams contains all the parameters to send to the API endpoint

	for the volumes create operation.

	Typically these are written to a http.Request.
*/
type VolumesCreateParams struct {

	/* AppName.

	   Fly App Name
	*/
	AppName string

	/* Request.

	   Request body
	*/
	Request *models.CreateVolumeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volumes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumesCreateParams) WithDefaults() *VolumesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volumes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volumes create params
func (o *VolumesCreateParams) WithTimeout(timeout time.Duration) *VolumesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volumes create params
func (o *VolumesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volumes create params
func (o *VolumesCreateParams) WithContext(ctx context.Context) *VolumesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volumes create params
func (o *VolumesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volumes create params
func (o *VolumesCreateParams) WithHTTPClient(client *http.Client) *VolumesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volumes create params
func (o *VolumesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the volumes create params
func (o *VolumesCreateParams) WithAppName(appName string) *VolumesCreateParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the volumes create params
func (o *VolumesCreateParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithRequest adds the request to the volumes create params
func (o *VolumesCreateParams) WithRequest(request *models.CreateVolumeRequest) *VolumesCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the volumes create params
func (o *VolumesCreateParams) SetRequest(request *models.CreateVolumeRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *VolumesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
