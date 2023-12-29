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

// AppsShowReader is a Reader for the AppsShow structure.
type AppsShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppsShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppsShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app_name}] Apps_show", response, response.Code())
	}
}

// NewAppsShowOK creates a AppsShowOK with default headers values
func NewAppsShowOK() *AppsShowOK {
	return &AppsShowOK{}
}

/*
AppsShowOK describes a response with status code 200, with default header values.

OK
*/
type AppsShowOK struct {
	Payload *models.App
}

// IsSuccess returns true when this apps show o k response has a 2xx status code
func (o *AppsShowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this apps show o k response has a 3xx status code
func (o *AppsShowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apps show o k response has a 4xx status code
func (o *AppsShowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this apps show o k response has a 5xx status code
func (o *AppsShowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this apps show o k response a status code equal to that given
func (o *AppsShowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the apps show o k response
func (o *AppsShowOK) Code() int {
	return 200
}

func (o *AppsShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_name}][%d] appsShowOK  %+v", 200, o.Payload)
}

func (o *AppsShowOK) String() string {
	return fmt.Sprintf("[GET /apps/{app_name}][%d] appsShowOK  %+v", 200, o.Payload)
}

func (o *AppsShowOK) GetPayload() *models.App {
	return o.Payload
}

func (o *AppsShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
