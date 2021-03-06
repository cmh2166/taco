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

// File file
// swagger:model File
type File struct {

	// UUID
	// Required: true
	UUID *string `json:"UUID"`

	// filename
	// Required: true
	Filename *string `json:"filename"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("UUID", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

func (m *File) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *File) UnmarshalBinary(b []byte) error {
	var res File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
