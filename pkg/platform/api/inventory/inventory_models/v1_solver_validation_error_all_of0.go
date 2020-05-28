// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SolverValidationErrorAllOf0 Solver Error
//
// An error response, for when resolving a recipe fails. The error may be permanent.
// swagger:model v1SolverValidationErrorAllOf0
type V1SolverValidationErrorAllOf0 struct {

	// When true sending the same request again may result in a different response, when false this error will always be returned for the same request
	IsTransient *bool `json:"is_transient,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this v1 solver validation error all of0
func (m *V1SolverValidationErrorAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SolverValidationErrorAllOf0) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SolverValidationErrorAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolverValidationErrorAllOf0) UnmarshalBinary(b []byte) error {
	var res V1SolverValidationErrorAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
