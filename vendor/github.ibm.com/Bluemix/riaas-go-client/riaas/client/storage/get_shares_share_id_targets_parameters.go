// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewGetSharesShareIDTargetsParams creates a new GetSharesShareIDTargetsParams object
// with the default values initialized.
func NewGetSharesShareIDTargetsParams() *GetSharesShareIDTargetsParams {
	var ()
	return &GetSharesShareIDTargetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSharesShareIDTargetsParamsWithTimeout creates a new GetSharesShareIDTargetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSharesShareIDTargetsParamsWithTimeout(timeout time.Duration) *GetSharesShareIDTargetsParams {
	var ()
	return &GetSharesShareIDTargetsParams{

		timeout: timeout,
	}
}

// NewGetSharesShareIDTargetsParamsWithContext creates a new GetSharesShareIDTargetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSharesShareIDTargetsParamsWithContext(ctx context.Context) *GetSharesShareIDTargetsParams {
	var ()
	return &GetSharesShareIDTargetsParams{

		Context: ctx,
	}
}

// NewGetSharesShareIDTargetsParamsWithHTTPClient creates a new GetSharesShareIDTargetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSharesShareIDTargetsParamsWithHTTPClient(client *http.Client) *GetSharesShareIDTargetsParams {
	var ()
	return &GetSharesShareIDTargetsParams{
		HTTPClient: client,
	}
}

/*GetSharesShareIDTargetsParams contains all the parameters to send to the API endpoint
for the get shares share ID targets operation typically these are written to a http.Request
*/
type GetSharesShareIDTargetsParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ResourceGroupID
	  Filters the collection to resources within one of the resource groups identified in a comma-separated list of resource group identifiers

	*/
	ResourceGroupID *string
	/*ShareID
	  The file share identifier

	*/
	ShareID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithTimeout(timeout time.Duration) *GetSharesShareIDTargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithContext(ctx context.Context) *GetSharesShareIDTargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithHTTPClient(client *http.Client) *GetSharesShareIDTargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithGeneration(generation int64) *GetSharesShareIDTargetsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithResourceGroupID adds the resourceGroupID to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithResourceGroupID(resourceGroupID *string) *GetSharesShareIDTargetsParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithShareID adds the shareID to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithShareID(shareID string) *GetSharesShareIDTargetsParams {
	o.SetShareID(shareID)
	return o
}

// SetShareID adds the shareId to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetShareID(shareID string) {
	o.ShareID = shareID
}

// WithVersion adds the version to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) WithVersion(version string) *GetSharesShareIDTargetsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get shares share ID targets params
func (o *GetSharesShareIDTargetsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSharesShareIDTargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	// path param share_id
	if err := r.SetPathParam("share_id", o.ShareID); err != nil {
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
