// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ConfigFileReader is a Reader for the ConfigFile structure.
type ConfigFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfigFileOK creates a ConfigFileOK with default headers values
func NewConfigFileOK() *ConfigFileOK {
	return &ConfigFileOK{}
}

/*ConfigFileOK handles this case with default header values.

Config File Download
*/
type ConfigFileOK struct {
	Payload string
}

func (o *ConfigFileOK) Error() string {
	return fmt.Sprintf("[GET /config][%d] configFileOK  %+v", 200, o.Payload)
}

func (o *ConfigFileOK) GetPayload() string {
	return o.Payload
}

func (o *ConfigFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
