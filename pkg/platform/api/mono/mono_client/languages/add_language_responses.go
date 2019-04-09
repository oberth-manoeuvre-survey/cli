// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// AddLanguageReader is a Reader for the AddLanguage structure.
type AddLanguageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLanguageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddLanguageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddLanguageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddLanguageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddLanguageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddLanguageOK creates a AddLanguageOK with default headers values
func NewAddLanguageOK() *AddLanguageOK {
	return &AddLanguageOK{}
}

/*AddLanguageOK handles this case with default header values.

Language Created
*/
type AddLanguageOK struct {
	Payload *mono_models.Language
}

func (o *AddLanguageOK) Error() string {
	return fmt.Sprintf("[POST /languages][%d] addLanguageOK  %+v", 200, o.Payload)
}

func (o *AddLanguageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Language)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLanguageBadRequest creates a AddLanguageBadRequest with default headers values
func NewAddLanguageBadRequest() *AddLanguageBadRequest {
	return &AddLanguageBadRequest{}
}

/*AddLanguageBadRequest handles this case with default header values.

Bad Request
*/
type AddLanguageBadRequest struct {
	Payload *mono_models.Message
}

func (o *AddLanguageBadRequest) Error() string {
	return fmt.Sprintf("[POST /languages][%d] addLanguageBadRequest  %+v", 400, o.Payload)
}

func (o *AddLanguageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLanguageForbidden creates a AddLanguageForbidden with default headers values
func NewAddLanguageForbidden() *AddLanguageForbidden {
	return &AddLanguageForbidden{}
}

/*AddLanguageForbidden handles this case with default header values.

Forbidden
*/
type AddLanguageForbidden struct {
	Payload *mono_models.Message
}

func (o *AddLanguageForbidden) Error() string {
	return fmt.Sprintf("[POST /languages][%d] addLanguageForbidden  %+v", 403, o.Payload)
}

func (o *AddLanguageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLanguageConflict creates a AddLanguageConflict with default headers values
func NewAddLanguageConflict() *AddLanguageConflict {
	return &AddLanguageConflict{}
}

/*AddLanguageConflict handles this case with default header values.

Conflict
*/
type AddLanguageConflict struct {
	Payload *mono_models.Message
}

func (o *AddLanguageConflict) Error() string {
	return fmt.Sprintf("[POST /languages][%d] addLanguageConflict  %+v", 409, o.Payload)
}

func (o *AddLanguageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}