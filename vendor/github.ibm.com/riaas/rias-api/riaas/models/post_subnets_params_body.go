// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostSubnetsParamsBody SubnetTemplate
// swagger:model postSubnetsParamsBody
type PostSubnetsParamsBody struct {

	// generation
	Generation Generation `json:"generation,omitempty"`

	// The IP version(s) supported by this subnet; if unspecified, `ipv4` is used
	// Enum: [ipv4 ipv6 both]
	IPVersion string `json:"ip_version,omitempty"`

	// The IPv4 range of the subnet, expressed in CIDR format
	IPV4CidrBlock string `json:"ipv4_cidr_block,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// network acl
	NetworkACL *PostSubnetsParamsBodyNetworkACL `json:"network_acl,omitempty"`

	// public gateway
	PublicGateway *PostSubnetsParamsBodyPublicGateway `json:"public_gateway,omitempty"`

	// resource group
	ResourceGroup *PostSubnetsParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// The total number of IPv4 addresses required
	TotalIPV4AddressCount int64 `json:"total_ipv4_address_count,omitempty"`

	// vpc
	Vpc *PostSubnetsParamsBodyVpc `json:"vpc,omitempty"`

	// zone
	Zone *NameReference `json:"zone,omitempty"`
}

// Validate validates this post subnets params body
func (m *PostSubnetsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSubnetsParamsBody) validateGeneration(formats strfmt.Registry) error {

	if swag.IsZero(m.Generation) { // not required
		return nil
	}

	if err := m.Generation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("generation")
		}
		return err
	}

	return nil
}

var postSubnetsParamsBodyTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postSubnetsParamsBodyTypeIPVersionPropEnum = append(postSubnetsParamsBodyTypeIPVersionPropEnum, v)
	}
}

const (

	// PostSubnetsParamsBodyIPVersionIPV4 captures enum value "ipv4"
	PostSubnetsParamsBodyIPVersionIPV4 string = "ipv4"

	// PostSubnetsParamsBodyIPVersionIPV6 captures enum value "ipv6"
	PostSubnetsParamsBodyIPVersionIPV6 string = "ipv6"

	// PostSubnetsParamsBodyIPVersionBoth captures enum value "both"
	PostSubnetsParamsBodyIPVersionBoth string = "both"
)

// prop value enum
func (m *PostSubnetsParamsBody) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postSubnetsParamsBodyTypeIPVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostSubnetsParamsBody) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPVersionEnum("ip_version", "body", m.IPVersion); err != nil {
		return err
	}

	return nil
}

func (m *PostSubnetsParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostSubnetsParamsBody) validateNetworkACL(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkACL) { // not required
		return nil
	}

	if m.NetworkACL != nil {
		if err := m.NetworkACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_acl")
			}
			return err
		}
	}

	return nil
}

func (m *PostSubnetsParamsBody) validatePublicGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicGateway) { // not required
		return nil
	}

	if m.PublicGateway != nil {
		if err := m.PublicGateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_gateway")
			}
			return err
		}
	}

	return nil
}

func (m *PostSubnetsParamsBody) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

func (m *PostSubnetsParamsBody) validateVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

func (m *PostSubnetsParamsBody) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSubnetsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSubnetsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostSubnetsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
