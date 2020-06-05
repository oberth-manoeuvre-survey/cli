// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GpuArchitectureCore GPU Architecture Core
//
// The properties of a GPU architecture needed to create a new one
//
// swagger:model v1GpuArchitectureCore
type V1GpuArchitectureCore struct {
	V1GpuArchitectureCoreAllOf0

	V1Revision
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1GpuArchitectureCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1GpuArchitectureCoreAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1GpuArchitectureCoreAllOf0 = aO0

	// AO1
	var aO1 V1Revision
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1Revision = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1GpuArchitectureCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1GpuArchitectureCoreAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1Revision)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 gpu architecture core
func (m *V1GpuArchitectureCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1GpuArchitectureCoreAllOf0
	if err := m.V1GpuArchitectureCoreAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1Revision
	if err := m.V1Revision.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1GpuArchitectureCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GpuArchitectureCore) UnmarshalBinary(b []byte) error {
	var res V1GpuArchitectureCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}