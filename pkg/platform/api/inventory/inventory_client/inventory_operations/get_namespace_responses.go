// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetNamespaceReader is a Reader for the GetNamespace structure.
type GetNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNamespaceOK creates a GetNamespaceOK with default headers values
func NewGetNamespaceOK() *GetNamespaceOK {
	return &GetNamespaceOK{}
}

/*GetNamespaceOK handles this case with default header values.

The retrieved namespace
*/
type GetNamespaceOK struct {
	Payload *inventory_models.V1Namespace
}

func (o *GetNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/{namespace}][%d] getNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceDefault creates a GetNamespaceDefault with default headers values
func NewGetNamespaceDefault(code int) *GetNamespaceDefault {
	return &GetNamespaceDefault{
		_statusCode: code,
	}
}

/*GetNamespaceDefault handles this case with default header values.

generic error response
*/
type GetNamespaceDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get namespace default response
func (o *GetNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *GetNamespaceDefault) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/{namespace}][%d] getNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *GetNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
