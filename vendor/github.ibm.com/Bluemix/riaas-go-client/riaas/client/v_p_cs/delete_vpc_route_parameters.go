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

// NewDeleteVpcRouteParams creates a new DeleteVpcRouteParams object
// with the default values initialized.
func NewDeleteVpcRouteParams() *DeleteVpcRouteParams {
	var ()
	return &DeleteVpcRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpcRouteParamsWithTimeout creates a new DeleteVpcRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpcRouteParamsWithTimeout(timeout time.Duration) *DeleteVpcRouteParams {
	var ()
	return &DeleteVpcRouteParams{

		timeout: timeout,
	}
}

// NewDeleteVpcRouteParamsWithContext creates a new DeleteVpcRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpcRouteParamsWithContext(ctx context.Context) *DeleteVpcRouteParams {
	var ()
	return &DeleteVpcRouteParams{

		Context: ctx,
	}
}

// NewDeleteVpcRouteParamsWithHTTPClient creates a new DeleteVpcRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpcRouteParamsWithHTTPClient(client *http.Client) *DeleteVpcRouteParams {
	var ()
	return &DeleteVpcRouteParams{
		HTTPClient: client,
	}
}

/*DeleteVpcRouteParams contains all the parameters to send to the API endpoint
for the delete vpc route operation typically these are written to a http.Request
*/
type DeleteVpcRouteParams struct {

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

// WithTimeout adds the timeout to the delete vpc route params
func (o *DeleteVpcRouteParams) WithTimeout(timeout time.Duration) *DeleteVpcRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vpc route params
func (o *DeleteVpcRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vpc route params
func (o *DeleteVpcRouteParams) WithContext(ctx context.Context) *DeleteVpcRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vpc route params
func (o *DeleteVpcRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vpc route params
func (o *DeleteVpcRouteParams) WithHTTPClient(client *http.Client) *DeleteVpcRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vpc route params
func (o *DeleteVpcRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete vpc route params
func (o *DeleteVpcRouteParams) WithGeneration(generation int64) *DeleteVpcRouteParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete vpc route params
func (o *DeleteVpcRouteParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete vpc route params
func (o *DeleteVpcRouteParams) WithID(id string) *DeleteVpcRouteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vpc route params
func (o *DeleteVpcRouteParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete vpc route params
func (o *DeleteVpcRouteParams) WithVersion(version string) *DeleteVpcRouteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete vpc route params
func (o *DeleteVpcRouteParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the delete vpc route params
func (o *DeleteVpcRouteParams) WithVpcID(vpcID string) *DeleteVpcRouteParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the delete vpc route params
func (o *DeleteVpcRouteParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpcRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
