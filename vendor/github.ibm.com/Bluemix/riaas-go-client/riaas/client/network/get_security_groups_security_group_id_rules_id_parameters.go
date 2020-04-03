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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSecurityGroupsSecurityGroupIDRulesIDParams creates a new GetSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized.
func NewGetSecurityGroupsSecurityGroupIDRulesIDParams() *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDRulesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout creates a new GetSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout(timeout time.Duration) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDRulesIDParams{

		timeout: timeout,
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithContext creates a new GetSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithContext(ctx context.Context) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDRulesIDParams{

		Context: ctx,
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithHTTPClient creates a new GetSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithHTTPClient(client *http.Client) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDRulesIDParams{
		HTTPClient: client,
	}
}

/*GetSecurityGroupsSecurityGroupIDRulesIDParams contains all the parameters to send to the API endpoint
for the get security groups security group ID rules ID operation typically these are written to a http.Request
*/
type GetSecurityGroupsSecurityGroupIDRulesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The rule identifier

	*/
	ID string
	/*SecurityGroupID
	  The security group identifier

	*/
	SecurityGroupID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithTimeout(timeout time.Duration) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithContext(ctx context.Context) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithHTTPClient(client *http.Client) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithGeneration(generation int64) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithID(id string) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetID(id string) {
	o.ID = id
}

// WithSecurityGroupID adds the securityGroupID to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithSecurityGroupID(securityGroupID string) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WithVersion adds the version to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WithVersion(version string) *GetSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get security groups security group ID rules ID params
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityGroupsSecurityGroupIDRulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param security_group_id
	if err := r.SetPathParam("security_group_id", o.SecurityGroupID); err != nil {
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
