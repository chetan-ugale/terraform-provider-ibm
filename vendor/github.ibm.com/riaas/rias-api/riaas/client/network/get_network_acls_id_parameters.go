// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetNetworkAclsIDParams creates a new GetNetworkAclsIDParams object
// with the default values initialized.
func NewGetNetworkAclsIDParams() *GetNetworkAclsIDParams {
	var ()
	return &GetNetworkAclsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsIDParamsWithTimeout creates a new GetNetworkAclsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsIDParamsWithTimeout(timeout time.Duration) *GetNetworkAclsIDParams {
	var ()
	return &GetNetworkAclsIDParams{

		timeout: timeout,
	}
}

// NewGetNetworkAclsIDParamsWithContext creates a new GetNetworkAclsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsIDParamsWithContext(ctx context.Context) *GetNetworkAclsIDParams {
	var ()
	return &GetNetworkAclsIDParams{

		Context: ctx,
	}
}

// NewGetNetworkAclsIDParamsWithHTTPClient creates a new GetNetworkAclsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsIDParamsWithHTTPClient(client *http.Client) *GetNetworkAclsIDParams {
	var ()
	return &GetNetworkAclsIDParams{
		HTTPClient: client,
	}
}

/*GetNetworkAclsIDParams contains all the parameters to send to the API endpoint
for the get network acls ID operation typically these are written to a http.Request
*/
type GetNetworkAclsIDParams struct {

	/*ID
	  The network ACL identifier

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

// WithTimeout adds the timeout to the get network acls ID params
func (o *GetNetworkAclsIDParams) WithTimeout(timeout time.Duration) *GetNetworkAclsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls ID params
func (o *GetNetworkAclsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls ID params
func (o *GetNetworkAclsIDParams) WithContext(ctx context.Context) *GetNetworkAclsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls ID params
func (o *GetNetworkAclsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls ID params
func (o *GetNetworkAclsIDParams) WithHTTPClient(client *http.Client) *GetNetworkAclsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls ID params
func (o *GetNetworkAclsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network acls ID params
func (o *GetNetworkAclsIDParams) WithID(id string) *GetNetworkAclsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network acls ID params
func (o *GetNetworkAclsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get network acls ID params
func (o *GetNetworkAclsIDParams) WithVersion(version string) *GetNetworkAclsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls ID params
func (o *GetNetworkAclsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
