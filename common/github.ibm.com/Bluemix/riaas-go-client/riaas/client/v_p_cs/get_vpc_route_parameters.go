// Code generated by go-swagger; DO NOT EDIT.

package v_p_cs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVpcRouteParams creates a new GetVpcRouteParams object
// with the default values initialized.
func NewGetVpcRouteParams() *GetVpcRouteParams {
	var ()
	return &GetVpcRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpcRouteParamsWithTimeout creates a new GetVpcRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpcRouteParamsWithTimeout(timeout time.Duration) *GetVpcRouteParams {
	var ()
	return &GetVpcRouteParams{

		timeout: timeout,
	}
}

// NewGetVpcRouteParamsWithContext creates a new GetVpcRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpcRouteParamsWithContext(ctx context.Context) *GetVpcRouteParams {
	var ()
	return &GetVpcRouteParams{

		Context: ctx,
	}
}

// NewGetVpcRouteParamsWithHTTPClient creates a new GetVpcRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpcRouteParamsWithHTTPClient(client *http.Client) *GetVpcRouteParams {
	var ()
	return &GetVpcRouteParams{
		HTTPClient: client,
	}
}

/*GetVpcRouteParams contains all the parameters to send to the API endpoint
for the get vpc route operation typically these are written to a http.Request
*/
type GetVpcRouteParams struct {

	/*Generation
	  The infrastructure generation for the request. For the API behavior documented here, use `1` or `2`.

	*/
	Generation int64
	/*ID
	  The route identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpc route params
func (o *GetVpcRouteParams) WithTimeout(timeout time.Duration) *GetVpcRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpc route params
func (o *GetVpcRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpc route params
func (o *GetVpcRouteParams) WithContext(ctx context.Context) *GetVpcRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpc route params
func (o *GetVpcRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpc route params
func (o *GetVpcRouteParams) WithHTTPClient(client *http.Client) *GetVpcRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpc route params
func (o *GetVpcRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpc route params
func (o *GetVpcRouteParams) WithGeneration(generation int64) *GetVpcRouteParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpc route params
func (o *GetVpcRouteParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get vpc route params
func (o *GetVpcRouteParams) WithID(id string) *GetVpcRouteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpc route params
func (o *GetVpcRouteParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get vpc route params
func (o *GetVpcRouteParams) WithVersion(version string) *GetVpcRouteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpc route params
func (o *GetVpcRouteParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the get vpc route params
func (o *GetVpcRouteParams) WithVpcID(vpcID string) *GetVpcRouteParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get vpc route params
func (o *GetVpcRouteParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpcRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

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

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
