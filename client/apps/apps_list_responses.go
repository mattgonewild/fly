// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattgonewild/fly/models"
)

// AppsListReader is a Reader for the AppsList structure.
type AppsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /apps] Apps_list", response, response.Code())
	}
}

// NewAppsListOK creates a AppsListOK with default headers values
func NewAppsListOK() *AppsListOK {
	return &AppsListOK{}
}

/*
AppsListOK describes a response with status code 200, with default header values.

OK
*/
type AppsListOK struct {
	Payload *models.ListAppsResponse
}

// IsSuccess returns true when this apps list o k response has a 2xx status code
func (o *AppsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this apps list o k response has a 3xx status code
func (o *AppsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apps list o k response has a 4xx status code
func (o *AppsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this apps list o k response has a 5xx status code
func (o *AppsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this apps list o k response a status code equal to that given
func (o *AppsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the apps list o k response
func (o *AppsListOK) Code() int {
	return 200
}

func (o *AppsListOK) Error() string {
	return fmt.Sprintf("[GET /apps][%d] appsListOK  %+v", 200, o.Payload)
}

func (o *AppsListOK) String() string {
	return fmt.Sprintf("[GET /apps][%d] appsListOK  %+v", 200, o.Payload)
}

func (o *AppsListOK) GetPayload() *models.ListAppsResponse {
	return o.Payload
}

func (o *AppsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAppsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
