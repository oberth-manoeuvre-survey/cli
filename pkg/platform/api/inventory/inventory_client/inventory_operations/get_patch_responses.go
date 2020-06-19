// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetPatchReader is a Reader for the GetPatch structure.
type GetPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPatchOK creates a GetPatchOK with default headers values
func NewGetPatchOK() *GetPatchOK {
	return &GetPatchOK{}
}

/*GetPatchOK handles this case with default header values.

The retrieved patch
*/
type GetPatchOK struct {
	Payload *inventory_models.V1Patch
}

func (o *GetPatchOK) Error() string {
	return fmt.Sprintf("[GET /v1/patches/{patch_id}][%d] getPatchOK  %+v", 200, o.Payload)
}

func (o *GetPatchOK) GetPayload() *inventory_models.V1Patch {
	return o.Payload
}

func (o *GetPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Patch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatchDefault creates a GetPatchDefault with default headers values
func NewGetPatchDefault(code int) *GetPatchDefault {
	return &GetPatchDefault{
		_statusCode: code,
	}
}

/*GetPatchDefault handles this case with default header values.

generic error response
*/
type GetPatchDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get patch default response
func (o *GetPatchDefault) Code() int {
	return o._statusCode
}

func (o *GetPatchDefault) Error() string {
	return fmt.Sprintf("[GET /v1/patches/{patch_id}][%d] getPatch default  %+v", o._statusCode, o.Payload)
}

func (o *GetPatchDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
