// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetSecurityGroupsSecurityGroupIDRulesReader is a Reader for the GetSecurityGroupsSecurityGroupIDRules structure.
type GetSecurityGroupsSecurityGroupIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsSecurityGroupIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsSecurityGroupIDRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSecurityGroupsSecurityGroupIDRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesOK creates a GetSecurityGroupsSecurityGroupIDRulesOK with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesOK() *GetSecurityGroupsSecurityGroupIDRulesOK {
	return &GetSecurityGroupsSecurityGroupIDRulesOK{}
}

/*GetSecurityGroupsSecurityGroupIDRulesOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsSecurityGroupIDRulesOK struct {
	Payload *models.GetSecurityGroupsSecurityGroupIDRulesOKBody
}

func (o *GetSecurityGroupsSecurityGroupIDRulesOK) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules][%d] getSecurityGroupsSecurityGroupIdRulesOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSecurityGroupsSecurityGroupIDRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsSecurityGroupIDRulesDefault creates a GetSecurityGroupsSecurityGroupIDRulesDefault with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesDefault(code int) *GetSecurityGroupsSecurityGroupIDRulesDefault {
	return &GetSecurityGroupsSecurityGroupIDRulesDefault{
		_statusCode: code,
	}
}

/*GetSecurityGroupsSecurityGroupIDRulesDefault handles this case with default header values.

unexpectederror
*/
type GetSecurityGroupsSecurityGroupIDRulesDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get security groups security group ID rules default response
func (o *GetSecurityGroupsSecurityGroupIDRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityGroupsSecurityGroupIDRulesDefault) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules][%d] GetSecurityGroupsSecurityGroupIDRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
