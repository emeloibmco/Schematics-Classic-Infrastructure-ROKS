// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageReference image reference
// swagger:model ImageReference
type ImageReference struct {

	// Creation Date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Description
	// Required: true
	Description *string `json:"description"`

	// Link to Image resource
	// Required: true
	Href *string `json:"href"`

	// Image ID
	// Required: true
	ImageID *string `json:"imageID"`

	// Last Update Date
	// Required: true
	// Format: date-time
	LastUpdateDate *strfmt.DateTime `json:"lastUpdateDate"`

	// Image Name
	// Required: true
	Name *string `json:"name"`

	// specifications
	// Required: true
	Specifications *ImageSpecifications `json:"specifications"`

	// Image State
	// Required: true
	State *string `json:"state"`

	// Storage type for image
	// Required: true
	StorageType *string `json:"storageType"`
}

// Validate validates this image reference
func (m *ImageReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageReference) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateSpecifications(formats strfmt.Registry) error {

	if err := validate.Required("specifications", "body", m.Specifications); err != nil {
		return err
	}

	if m.Specifications != nil {
		if err := m.Specifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specifications")
			}
			return err
		}
	}

	return nil
}

func (m *ImageReference) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateStorageType(formats strfmt.Registry) error {

	if err := validate.Required("storageType", "body", m.StorageType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageReference) UnmarshalBinary(b []byte) error {
	var res ImageReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
