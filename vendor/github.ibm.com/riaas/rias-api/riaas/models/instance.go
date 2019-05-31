// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Instance instance
// swagger:model instance
type Instance struct {

	// boot volume attachment
	BootVolumeAttachment *VolumeAttachmentReference `json:"boot_volume_attachment,omitempty"`

	// cpu
	CPU *InstanceCPU `json:"cpu,omitempty"`

	// The date and time that the instance was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this instance
	Crn string `json:"crn,omitempty"`

	// generation
	Generation Generation `json:"generation,omitempty"`

	// gpu
	Gpu *InstanceGpu `json:"gpu,omitempty"`

	// The URL for this instance
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this instance
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// image
	Image *ResourceReference `json:"image,omitempty"`

	// The amount of memory in megabytes
	// Maximum: 524288
	// Minimum: 512
	// Multiple Of: 256
	Memory int64 `json:"memory,omitempty"`

	// The user-defined name for this instance
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// Collection of the instance's network interfaces, not including the primary network interface
	NetworkInterfaces []*NetworkInterfaceReference `json:"network_interfaces"`

	// primary network interface
	PrimaryNetworkInterface *NetworkInterfaceReference `json:"primary_network_interface,omitempty"`

	// profile
	Profile *NameReference `json:"profile,omitempty"`

	// resource group
	ResourceGroup *GroupReference `json:"resource_group,omitempty"`

	// The status of the instance
	// Enum: [stopped starting running pausing paused resuming stopping restarting]
	Status string `json:"status,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// instance Type
	// Enum: [virtual]
	Type string `json:"type,omitempty"`

	// Collection of volume interfaces
	VolumeAttachments []*VolumeAttachmentReference `json:"volume_attachments,omitempty"`

	// vpc
	Vpc *ResourceReference `json:"vpc,omitempty"`

	// zone
	Zone *ZoneReference `json:"zone,omitempty"`
}

// Validate validates this instance
func (m *Instance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootVolumeAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNetworkInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeAttachments(formats); err != nil {
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

func (m *Instance) validateBootVolumeAttachment(formats strfmt.Registry) error {

	if swag.IsZero(m.BootVolumeAttachment) { // not required
		return nil
	}

	if m.BootVolumeAttachment != nil {
		if err := m.BootVolumeAttachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boot_volume_attachment")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateGeneration(formats strfmt.Registry) error {

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

func (m *Instance) validateGpu(formats strfmt.Registry) error {

	if swag.IsZero(m.Gpu) { // not required
		return nil
	}

	if m.Gpu != nil {
		if err := m.Gpu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpu")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory", "body", int64(m.Memory), 512, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory", "body", int64(m.Memory), 524288, false); err != nil {
		return err
	}

	if err := validate.MultipleOf("memory", "body", float64(m.Memory), 256); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateNetworkInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInterfaces); i++ {
		if swag.IsZero(m.NetworkInterfaces[i]) { // not required
			continue
		}

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Instance) validatePrimaryNetworkInterface(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryNetworkInterface) { // not required
		return nil
	}

	if m.PrimaryNetworkInterface != nil {
		if err := m.PrimaryNetworkInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_network_interface")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateResourceGroup(formats strfmt.Registry) error {

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

var instanceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["stopped","starting","running","pausing","paused","resuming","stopping","restarting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeStatusPropEnum = append(instanceTypeStatusPropEnum, v)
	}
}

const (

	// InstanceStatusStopped captures enum value "stopped"
	InstanceStatusStopped string = "stopped"

	// InstanceStatusStarting captures enum value "starting"
	InstanceStatusStarting string = "starting"

	// InstanceStatusRunning captures enum value "running"
	InstanceStatusRunning string = "running"

	// InstanceStatusPausing captures enum value "pausing"
	InstanceStatusPausing string = "pausing"

	// InstanceStatusPaused captures enum value "paused"
	InstanceStatusPaused string = "paused"

	// InstanceStatusResuming captures enum value "resuming"
	InstanceStatusResuming string = "resuming"

	// InstanceStatusStopping captures enum value "stopping"
	InstanceStatusStopping string = "stopping"

	// InstanceStatusRestarting captures enum value "restarting"
	InstanceStatusRestarting string = "restarting"
)

// prop value enum
func (m *Instance) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instance) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var instanceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeTypePropEnum = append(instanceTypeTypePropEnum, v)
	}
}

const (

	// InstanceTypeVirtual captures enum value "virtual"
	InstanceTypeVirtual string = "virtual"
)

// prop value enum
func (m *Instance) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instance) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateVolumeAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeAttachments) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeAttachments); i++ {
		if swag.IsZero(m.VolumeAttachments[i]) { // not required
			continue
		}

		if m.VolumeAttachments[i] != nil {
			if err := m.VolumeAttachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Instance) validateVpc(formats strfmt.Registry) error {

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

func (m *Instance) validateZone(formats strfmt.Registry) error {

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
func (m *Instance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instance) UnmarshalBinary(b []byte) error {
	var res Instance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
