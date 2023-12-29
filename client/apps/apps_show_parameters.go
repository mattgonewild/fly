// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewAppsShowParams creates a new AppsShowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppsShowParams() *AppsShowParams {
	return &AppsShowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppsShowParamsWithTimeout creates a new AppsShowParams object
// with the ability to set a timeout on a request.
func NewAppsShowParamsWithTimeout(timeout time.Duration) *AppsShowParams {
	return &AppsShowParams{
		timeout: timeout,
	}
}

// NewAppsShowParamsWithContext creates a new AppsShowParams object
// with the ability to set a context for a request.
func NewAppsShowParamsWithContext(ctx context.Context) *AppsShowParams {
	return &AppsShowParams{
		Context: ctx,
	}
}

// NewAppsShowParamsWithHTTPClient creates a new AppsShowParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppsShowParamsWithHTTPClient(client *http.Client) *AppsShowParams {
	return &AppsShowParams{
		HTTPClient: client,
	}
}

/*
AppsShowParams contains all the parameters to send to the API endpoint

	for the apps show operation.

	Typically these are written to a http.Request.
*/
type AppsShowParams struct {

	/* AppName.

	   Fly App Name
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the apps show params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppsShowParams) WithDefaults() *AppsShowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the apps show params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppsShowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the apps show params
func (o *AppsShowParams) WithTimeout(timeout time.Duration) *AppsShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apps show params
func (o *AppsShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apps show params
func (o *AppsShowParams) WithContext(ctx context.Context) *AppsShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apps show params
func (o *AppsShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apps show params
func (o *AppsShowParams) WithHTTPClient(client *http.Client) *AppsShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apps show params
func (o *AppsShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the apps show params
func (o *AppsShowParams) WithAppName(appName string) *AppsShowParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the apps show params
func (o *AppsShowParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *AppsShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}