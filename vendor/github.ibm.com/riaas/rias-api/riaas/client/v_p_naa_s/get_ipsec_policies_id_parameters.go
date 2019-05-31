// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIpsecPoliciesIDParams creates a new GetIpsecPoliciesIDParams object
// with the default values initialized.
func NewGetIpsecPoliciesIDParams() *GetIpsecPoliciesIDParams {
	var ()
	return &GetIpsecPoliciesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpsecPoliciesIDParamsWithTimeout creates a new GetIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIpsecPoliciesIDParamsWithTimeout(timeout time.Duration) *GetIpsecPoliciesIDParams {
	var ()
	return &GetIpsecPoliciesIDParams{

		timeout: timeout,
	}
}

// NewGetIpsecPoliciesIDParamsWithContext creates a new GetIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIpsecPoliciesIDParamsWithContext(ctx context.Context) *GetIpsecPoliciesIDParams {
	var ()
	return &GetIpsecPoliciesIDParams{

		Context: ctx,
	}
}

// NewGetIpsecPoliciesIDParamsWithHTTPClient creates a new GetIpsecPoliciesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIpsecPoliciesIDParamsWithHTTPClient(client *http.Client) *GetIpsecPoliciesIDParams {
	var ()
	return &GetIpsecPoliciesIDParams{
		HTTPClient: client,
	}
}

/*GetIpsecPoliciesIDParams contains all the parameters to send to the API endpoint
for the get ipsec policies ID operation typically these are written to a http.Request
*/
type GetIpsecPoliciesIDParams struct {

	/*ID
	  The IPsec policy idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) WithTimeout(timeout time.Duration) *GetIpsecPoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) WithContext(ctx context.Context) *GetIpsecPoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) WithHTTPClient(client *http.Client) *GetIpsecPoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) WithID(id string) *GetIpsecPoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) WithVersion(version string) *GetIpsecPoliciesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get ipsec policies ID params
func (o *GetIpsecPoliciesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpsecPoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
