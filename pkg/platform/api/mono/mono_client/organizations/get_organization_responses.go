// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetOrganizationReader is a Reader for the GetOrganization structure.
type GetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationOK creates a GetOrganizationOK with default headers values
func NewGetOrganizationOK() *GetOrganizationOK {
	return &GetOrganizationOK{}
}

/*GetOrganizationOK handles this case with default header values.

Organization Record
*/
type GetOrganizationOK struct {
	Payload *mono_models.Organization
}

func (o *GetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}][%d] getOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationNotFound creates a GetOrganizationNotFound with default headers values
func NewGetOrganizationNotFound() *GetOrganizationNotFound {
	return &GetOrganizationNotFound{}
}

/*GetOrganizationNotFound handles this case with default header values.

Not Found
*/
type GetOrganizationNotFound struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}][%d] getOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInternalServerError creates a GetOrganizationInternalServerError with default headers values
func NewGetOrganizationInternalServerError() *GetOrganizationInternalServerError {
	return &GetOrganizationInternalServerError{}
}

/*GetOrganizationInternalServerError handles this case with default header values.

Server Error
*/
type GetOrganizationInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}][%d] getOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
