// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// DeleteLoadBalancersIDListenersListenerIDReader is a Reader for the DeleteLoadBalancersIDListenersListenerID structure.
type DeleteLoadBalancersIDListenersListenerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancersIDListenersListenerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLoadBalancersIDListenersListenerIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteLoadBalancersIDListenersListenerIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancersIDListenersListenerIDNoContent creates a DeleteLoadBalancersIDListenersListenerIDNoContent with default headers values
func NewDeleteLoadBalancersIDListenersListenerIDNoContent() *DeleteLoadBalancersIDListenersListenerIDNoContent {
	return &DeleteLoadBalancersIDListenersListenerIDNoContent{}
}

/*DeleteLoadBalancersIDListenersListenerIDNoContent handles this case with default header values.

The listener deletion request was accepted.
*/
type DeleteLoadBalancersIDListenersListenerIDNoContent struct {
}

func (o *DeleteLoadBalancersIDListenersListenerIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}/listeners/{listener_id}][%d] deleteLoadBalancersIdListenersListenerIdNoContent ", 204)
}

func (o *DeleteLoadBalancersIDListenersListenerIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancersIDListenersListenerIDNotFound creates a DeleteLoadBalancersIDListenersListenerIDNotFound with default headers values
func NewDeleteLoadBalancersIDListenersListenerIDNotFound() *DeleteLoadBalancersIDListenersListenerIDNotFound {
	return &DeleteLoadBalancersIDListenersListenerIDNotFound{}
}

/*DeleteLoadBalancersIDListenersListenerIDNotFound handles this case with default header values.

A load balancer or listener with the specified identifier could not be found.
*/
type DeleteLoadBalancersIDListenersListenerIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteLoadBalancersIDListenersListenerIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}/listeners/{listener_id}][%d] deleteLoadBalancersIdListenersListenerIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLoadBalancersIDListenersListenerIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
