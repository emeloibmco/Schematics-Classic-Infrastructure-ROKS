// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoadBalancerCollection load balancer collection
// swagger:model LoadBalancerCollection
type LoadBalancerCollection struct {

	// first
	First *LoadBalancerCollectionFirst `json:"first,omitempty"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// Collection of load balancers
	LoadBalancers []*LoadBalancer `json:"load_balancers"`

	// next
	Next *LoadBalancerCollectionNext `json:"next,omitempty"`
}

// Validate validates this load balancer collection
func (m *LoadBalancerCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerCollection) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerCollection) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", int64(m.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancerCollection) validateLoadBalancers(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalancers) { // not required
		return nil
	}

	for i := 0; i < len(m.LoadBalancers); i++ {
		if swag.IsZero(m.LoadBalancers[i]) { // not required
			continue
		}

		if m.LoadBalancers[i] != nil {
			if err := m.LoadBalancers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("load_balancers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadBalancerCollection) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerCollection) UnmarshalBinary(b []byte) error {
	var res LoadBalancerCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadBalancerCollectionFirst load balancer collection first
// swagger:model LoadBalancerCollectionFirst
type LoadBalancerCollectionFirst struct {

	// href
	Href interface{} `json:"href,omitempty"`
}

// Validate validates this load balancer collection first
func (m *LoadBalancerCollectionFirst) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerCollectionFirst) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerCollectionFirst) UnmarshalBinary(b []byte) error {
	var res LoadBalancerCollectionFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadBalancerCollectionNext A reference to the next page of resources; this reference is included for all pages except the last page
// swagger:model LoadBalancerCollectionNext
type LoadBalancerCollectionNext struct {

	// The URL for the next page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this load balancer collection next
func (m *LoadBalancerCollectionNext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerCollectionNext) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("next"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	if err := validate.Pattern("next"+"."+"href", "body", string(*m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerCollectionNext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerCollectionNext) UnmarshalBinary(b []byte) error {
	var res LoadBalancerCollectionNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
