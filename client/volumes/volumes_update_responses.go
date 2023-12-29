// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattgonewild/fly/models"
)

// VolumesUpdateReader is a Reader for the VolumesUpdate structure.
type VolumesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVolumesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /apps/{app_name}/volumes/{volume_id}] Volumes_update", response, response.Code())
	}
}

// NewVolumesUpdateOK creates a VolumesUpdateOK with default headers values
func NewVolumesUpdateOK() *VolumesUpdateOK {
	return &VolumesUpdateOK{}
}

/*
VolumesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type VolumesUpdateOK struct {
	Payload *models.Volume
}

// IsSuccess returns true when this volumes update o k response has a 2xx status code
func (o *VolumesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volumes update o k response has a 3xx status code
func (o *VolumesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volumes update o k response has a 4xx status code
func (o *VolumesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volumes update o k response has a 5xx status code
func (o *VolumesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volumes update o k response a status code equal to that given
func (o *VolumesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volumes update o k response
func (o *VolumesUpdateOK) Code() int {
	return 200
}

func (o *VolumesUpdateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/volumes/{volume_id}][%d] volumesUpdateOK  %+v", 200, o.Payload)
}

func (o *VolumesUpdateOK) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/volumes/{volume_id}][%d] volumesUpdateOK  %+v", 200, o.Payload)
}

func (o *VolumesUpdateOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *VolumesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumesUpdateBadRequest creates a VolumesUpdateBadRequest with default headers values
func NewVolumesUpdateBadRequest() *VolumesUpdateBadRequest {
	return &VolumesUpdateBadRequest{}
}

/*
VolumesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VolumesUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volumes update bad request response has a 2xx status code
func (o *VolumesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volumes update bad request response has a 3xx status code
func (o *VolumesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volumes update bad request response has a 4xx status code
func (o *VolumesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this volumes update bad request response has a 5xx status code
func (o *VolumesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this volumes update bad request response a status code equal to that given
func (o *VolumesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the volumes update bad request response
func (o *VolumesUpdateBadRequest) Code() int {
	return 400
}

func (o *VolumesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/volumes/{volume_id}][%d] volumesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VolumesUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/volumes/{volume_id}][%d] volumesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VolumesUpdateBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
