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

// NewAppsListParams creates a new AppsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppsListParams() *AppsListParams {
	return &AppsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppsListParamsWithTimeout creates a new AppsListParams object
// with the ability to set a timeout on a request.
func NewAppsListParamsWithTimeout(timeout time.Duration) *AppsListParams {
	return &AppsListParams{
		timeout: timeout,
	}
}

// NewAppsListParamsWithContext creates a new AppsListParams object
// with the ability to set a context for a request.
func NewAppsListParamsWithContext(ctx context.Context) *AppsListParams {
	return &AppsListParams{
		Context: ctx,
	}
}

// NewAppsListParamsWithHTTPClient creates a new AppsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppsListParamsWithHTTPClient(client *http.Client) *AppsListParams {
	return &AppsListParams{
		HTTPClient: client,
	}
}

/*
AppsListParams contains all the parameters to send to the API endpoint

	for the apps list operation.

	Typically these are written to a http.Request.
*/
type AppsListParams struct {

	/* OrgSlug.

	   The org slug, or 'personal', to filter apps
	*/
	OrgSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the apps list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppsListParams) WithDefaults() *AppsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the apps list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the apps list params
func (o *AppsListParams) WithTimeout(timeout time.Duration) *AppsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apps list params
func (o *AppsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apps list params
func (o *AppsListParams) WithContext(ctx context.Context) *AppsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apps list params
func (o *AppsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apps list params
func (o *AppsListParams) WithHTTPClient(client *http.Client) *AppsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apps list params
func (o *AppsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgSlug adds the orgSlug to the apps list params
func (o *AppsListParams) WithOrgSlug(orgSlug string) *AppsListParams {
	o.SetOrgSlug(orgSlug)
	return o
}

// SetOrgSlug adds the orgSlug to the apps list params
func (o *AppsListParams) SetOrgSlug(orgSlug string) {
	o.OrgSlug = orgSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AppsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param org_slug
	qrOrgSlug := o.OrgSlug
	qOrgSlug := qrOrgSlug
	if qOrgSlug != "" {

		if err := r.SetQueryParam("org_slug", qOrgSlug); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
