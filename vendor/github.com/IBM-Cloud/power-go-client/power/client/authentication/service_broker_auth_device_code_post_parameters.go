// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewServiceBrokerAuthDeviceCodePostParams creates a new ServiceBrokerAuthDeviceCodePostParams object
// with the default values initialized.
func NewServiceBrokerAuthDeviceCodePostParams() *ServiceBrokerAuthDeviceCodePostParams {

	return &ServiceBrokerAuthDeviceCodePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerAuthDeviceCodePostParamsWithTimeout creates a new ServiceBrokerAuthDeviceCodePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerAuthDeviceCodePostParamsWithTimeout(timeout time.Duration) *ServiceBrokerAuthDeviceCodePostParams {

	return &ServiceBrokerAuthDeviceCodePostParams{

		timeout: timeout,
	}
}

// NewServiceBrokerAuthDeviceCodePostParamsWithContext creates a new ServiceBrokerAuthDeviceCodePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerAuthDeviceCodePostParamsWithContext(ctx context.Context) *ServiceBrokerAuthDeviceCodePostParams {

	return &ServiceBrokerAuthDeviceCodePostParams{

		Context: ctx,
	}
}

// NewServiceBrokerAuthDeviceCodePostParamsWithHTTPClient creates a new ServiceBrokerAuthDeviceCodePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerAuthDeviceCodePostParamsWithHTTPClient(client *http.Client) *ServiceBrokerAuthDeviceCodePostParams {

	return &ServiceBrokerAuthDeviceCodePostParams{
		HTTPClient: client,
	}
}

/*ServiceBrokerAuthDeviceCodePostParams contains all the parameters to send to the API endpoint
for the service broker auth device code post operation typically these are written to a http.Request
*/
type ServiceBrokerAuthDeviceCodePostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) WithTimeout(timeout time.Duration) *ServiceBrokerAuthDeviceCodePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) WithContext(ctx context.Context) *ServiceBrokerAuthDeviceCodePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) WithHTTPClient(client *http.Client) *ServiceBrokerAuthDeviceCodePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker auth device code post params
func (o *ServiceBrokerAuthDeviceCodePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerAuthDeviceCodePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
