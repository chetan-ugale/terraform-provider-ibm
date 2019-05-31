// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PatchKeysIDReader is a Reader for the PatchKeysID structure.
type PatchKeysIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKeysIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchKeysIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchKeysIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchKeysIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchKeysIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPatchKeysIDOK creates a PatchKeysIDOK with default headers values
func NewPatchKeysIDOK() *PatchKeysIDOK {
	return &PatchKeysIDOK{}
}

/*PatchKeysIDOK handles this case with default header values.

dummy
*/
type PatchKeysIDOK struct {
	Payload *models.Key
}

func (o *PatchKeysIDOK) Error() string {
	return fmt.Sprintf("[PATCH /keys/{id}][%d] patchKeysIdOK  %+v", 200, o.Payload)
}

func (o *PatchKeysIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Key)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKeysIDBadRequest creates a PatchKeysIDBadRequest with default headers values
func NewPatchKeysIDBadRequest() *PatchKeysIDBadRequest {
	return &PatchKeysIDBadRequest{}
}

/*PatchKeysIDBadRequest handles this case with default header values.

error
*/
type PatchKeysIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchKeysIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /keys/{id}][%d] patchKeysIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKeysIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKeysIDNotFound creates a PatchKeysIDNotFound with default headers values
func NewPatchKeysIDNotFound() *PatchKeysIDNotFound {
	return &PatchKeysIDNotFound{}
}

/*PatchKeysIDNotFound handles this case with default header values.

error
*/
type PatchKeysIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchKeysIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /keys/{id}][%d] patchKeysIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchKeysIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKeysIDDefault creates a PatchKeysIDDefault with default headers values
func NewPatchKeysIDDefault(code int) *PatchKeysIDDefault {
	return &PatchKeysIDDefault{
		_statusCode: code,
	}
}

/*PatchKeysIDDefault handles this case with default header values.

unexpectederror
*/
type PatchKeysIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the patch keys ID default response
func (o *PatchKeysIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchKeysIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /keys/{id}][%d] PatchKeysID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchKeysIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
