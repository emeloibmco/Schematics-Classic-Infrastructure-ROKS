// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PutLoadBalancersIDPoolsPoolIDMembersReader is a Reader for the PutLoadBalancersIDPoolsPoolIDMembers structure.
type PutLoadBalancersIDPoolsPoolIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLoadBalancersIDPoolsPoolIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPutLoadBalancersIDPoolsPoolIDMembersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutLoadBalancersIDPoolsPoolIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutLoadBalancersIDPoolsPoolIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLoadBalancersIDPoolsPoolIDMembersAccepted creates a PutLoadBalancersIDPoolsPoolIDMembersAccepted with default headers values
func NewPutLoadBalancersIDPoolsPoolIDMembersAccepted() *PutLoadBalancersIDPoolsPoolIDMembersAccepted {
	return &PutLoadBalancersIDPoolsPoolIDMembersAccepted{}
}

/*PutLoadBalancersIDPoolsPoolIDMembersAccepted handles this case with default header values.

The update request was accepted.
*/
type PutLoadBalancersIDPoolsPoolIDMembersAccepted struct {
	Payload *models.MemberCollection
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersAccepted) Error() string {
	return fmt.Sprintf("[PUT /load_balancers/{id}/pools/{pool_id}/members][%d] putLoadBalancersIdPoolsPoolIdMembersAccepted  %+v", 202, o.Payload)
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLoadBalancersIDPoolsPoolIDMembersBadRequest creates a PutLoadBalancersIDPoolsPoolIDMembersBadRequest with default headers values
func NewPutLoadBalancersIDPoolsPoolIDMembersBadRequest() *PutLoadBalancersIDPoolsPoolIDMembersBadRequest {
	return &PutLoadBalancersIDPoolsPoolIDMembersBadRequest{}
}

/*PutLoadBalancersIDPoolsPoolIDMembersBadRequest handles this case with default header values.

An invalid member template was provided.
*/
type PutLoadBalancersIDPoolsPoolIDMembersBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[PUT /load_balancers/{id}/pools/{pool_id}/members][%d] putLoadBalancersIdPoolsPoolIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLoadBalancersIDPoolsPoolIDMembersNotFound creates a PutLoadBalancersIDPoolsPoolIDMembersNotFound with default headers values
func NewPutLoadBalancersIDPoolsPoolIDMembersNotFound() *PutLoadBalancersIDPoolsPoolIDMembersNotFound {
	return &PutLoadBalancersIDPoolsPoolIDMembersNotFound{}
}

/*PutLoadBalancersIDPoolsPoolIDMembersNotFound handles this case with default header values.

A load balancer or pool with the specified identifier could not be found.
*/
type PutLoadBalancersIDPoolsPoolIDMembersNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersNotFound) Error() string {
	return fmt.Sprintf("[PUT /load_balancers/{id}/pools/{pool_id}/members][%d] putLoadBalancersIdPoolsPoolIdMembersNotFound  %+v", 404, o.Payload)
}

func (o *PutLoadBalancersIDPoolsPoolIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
