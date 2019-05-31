// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetVpcsVpcIDAddressPrefixesOKBody AddressPoolPrefixCollection
// swagger:model getVpcsVpcIdAddressPrefixesOKBody
type GetVpcsVpcIDAddressPrefixesOKBody struct {

	// Collection of address prefixes
	AddressPrefixes []*AddressPoolPrefix `json:"address_prefixes,omitempty"`
}

// Validate validates this get vpcs vpc Id address prefixes o k body
func (m *GetVpcsVpcIDAddressPrefixesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressPrefixes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetVpcsVpcIDAddressPrefixesOKBody) validateAddressPrefixes(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressPrefixes) { // not required
		return nil
	}

	for i := 0; i < len(m.AddressPrefixes); i++ {
		if swag.IsZero(m.AddressPrefixes[i]) { // not required
			continue
		}

		if m.AddressPrefixes[i] != nil {
			if err := m.AddressPrefixes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("address_prefixes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetVpcsVpcIDAddressPrefixesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVpcsVpcIDAddressPrefixesOKBody) UnmarshalBinary(b []byte) error {
	var res GetVpcsVpcIDAddressPrefixesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
