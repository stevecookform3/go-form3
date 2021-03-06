// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ChargesInformation charges information
// swagger:model ChargesInformation
type ChargesInformation struct {

	// bearer code
	BearerCode string `json:"bearer_code,omitempty"`

	// receiver charges amount
	ReceiverChargesAmount string `json:"receiver_charges_amount,omitempty"`

	// receiver charges currency
	ReceiverChargesCurrency string `json:"receiver_charges_currency,omitempty"`

	// sender charges
	SenderCharges ChargesInformationSenderCharges `json:"sender_charges"`
}

// Validate validates this charges information
func (m *ChargesInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ChargesInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChargesInformation) UnmarshalBinary(b []byte) error {
	var res ChargesInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
