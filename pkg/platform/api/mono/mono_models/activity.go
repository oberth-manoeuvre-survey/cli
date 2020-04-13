// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Activity activity
//
// swagger:model Activity
type Activity struct {

	// scan
	Scan *Scan `json:"scan,omitempty"`

	// session
	Session *Session `json:"session,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this activity
func (m *Activity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Activity) validateScan(formats strfmt.Registry) error {

	if swag.IsZero(m.Scan) { // not required
		return nil
	}

	if m.Scan != nil {
		if err := m.Scan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan")
			}
			return err
		}
	}

	return nil
}

func (m *Activity) validateSession(formats strfmt.Registry) error {

	if swag.IsZero(m.Session) { // not required
		return nil
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

func (m *Activity) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Activity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Activity) UnmarshalBinary(b []byte) error {
	var res Activity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
