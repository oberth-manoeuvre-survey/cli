// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusMessageEnvelopeBody4 Task Failed
//
// A message indicating that a requested task for a build failed.
// swagger:model statusMessageEnvelopeBody4
type StatusMessageEnvelopeBody4 struct {

	// All of the errors from the failed task.
	// Required: true
	Errors []string `json:"errors"`

	// task
	// Required: true
	Task *string `json:"task"`
}

// Validate validates this status message envelope body4
func (m *StatusMessageEnvelopeBody4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusMessageEnvelopeBody4) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody4) validateTask(formats strfmt.Registry) error {

	if err := validate.Required("task", "body", m.Task); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody4) UnmarshalBinary(b []byte) error {
	var res StatusMessageEnvelopeBody4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
