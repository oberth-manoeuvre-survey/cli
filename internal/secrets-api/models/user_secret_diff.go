// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSecretDiff user secret diff
// swagger:model UserSecretDiff
type UserSecretDiff struct {

	// public key
	// Required: true
	PublicKey *string `json:"public_key"`

	// shares
	Shares []*UserSecretShare `json:"shares"`

	// user id
	// Required: true
	// Format: uuid
	UserID *strfmt.UUID `json:"user_id"`
}

// Validate validates this user secret diff
func (m *UserSecretDiff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSecretDiff) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *UserSecretDiff) validateShares(formats strfmt.Registry) error {

	if swag.IsZero(m.Shares) { // not required
		return nil
	}

	for i := 0; i < len(m.Shares); i++ {
		if swag.IsZero(m.Shares[i]) { // not required
			continue
		}

		if m.Shares[i] != nil {
			if err := m.Shares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserSecretDiff) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSecretDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSecretDiff) UnmarshalBinary(b []byte) error {
	var res UserSecretDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}