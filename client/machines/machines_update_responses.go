// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattgonewild/fly/models"
)

// MachinesUpdateReader is a Reader for the MachinesUpdate structure.
type MachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMachinesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /apps/{app_name}/machines/{machine_id}] Machines_update", response, response.Code())
	}
}

// NewMachinesUpdateOK creates a MachinesUpdateOK with default headers values
func NewMachinesUpdateOK() *MachinesUpdateOK {
	return &MachinesUpdateOK{}
}

/*
MachinesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type MachinesUpdateOK struct {
	Payload *models.Machine
}

// IsSuccess returns true when this machines update o k response has a 2xx status code
func (o *MachinesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machines update o k response has a 3xx status code
func (o *MachinesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machines update o k response has a 4xx status code
func (o *MachinesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machines update o k response has a 5xx status code
func (o *MachinesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machines update o k response a status code equal to that given
func (o *MachinesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machines update o k response
func (o *MachinesUpdateOK) Code() int {
	return 200
}

func (o *MachinesUpdateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}][%d] machinesUpdateOK  %+v", 200, o.Payload)
}

func (o *MachinesUpdateOK) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}][%d] machinesUpdateOK  %+v", 200, o.Payload)
}

func (o *MachinesUpdateOK) GetPayload() *models.Machine {
	return o.Payload
}

func (o *MachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Machine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachinesUpdateBadRequest creates a MachinesUpdateBadRequest with default headers values
func NewMachinesUpdateBadRequest() *MachinesUpdateBadRequest {
	return &MachinesUpdateBadRequest{}
}

/*
MachinesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MachinesUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this machines update bad request response has a 2xx status code
func (o *MachinesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this machines update bad request response has a 3xx status code
func (o *MachinesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machines update bad request response has a 4xx status code
func (o *MachinesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this machines update bad request response has a 5xx status code
func (o *MachinesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this machines update bad request response a status code equal to that given
func (o *MachinesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the machines update bad request response
func (o *MachinesUpdateBadRequest) Code() int {
	return 400
}

func (o *MachinesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}][%d] machinesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *MachinesUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}][%d] machinesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *MachinesUpdateBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MachinesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
