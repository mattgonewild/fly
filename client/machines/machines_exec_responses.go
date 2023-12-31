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

// MachinesExecReader is a Reader for the MachinesExec structure.
type MachinesExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachinesExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachinesExecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMachinesExecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /apps/{app_name}/machines/{machine_id}/exec] Machines_exec", response, response.Code())
	}
}

// NewMachinesExecOK creates a MachinesExecOK with default headers values
func NewMachinesExecOK() *MachinesExecOK {
	return &MachinesExecOK{}
}

/*
MachinesExecOK describes a response with status code 200, with default header values.

Raw command output bytes are written back
*/
type MachinesExecOK struct {
	Payload string
}

// IsSuccess returns true when this machines exec o k response has a 2xx status code
func (o *MachinesExecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machines exec o k response has a 3xx status code
func (o *MachinesExecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machines exec o k response has a 4xx status code
func (o *MachinesExecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machines exec o k response has a 5xx status code
func (o *MachinesExecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machines exec o k response a status code equal to that given
func (o *MachinesExecOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machines exec o k response
func (o *MachinesExecOK) Code() int {
	return 200
}

func (o *MachinesExecOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}/exec][%d] machinesExecOK  %+v", 200, o.Payload)
}

func (o *MachinesExecOK) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}/exec][%d] machinesExecOK  %+v", 200, o.Payload)
}

func (o *MachinesExecOK) GetPayload() string {
	return o.Payload
}

func (o *MachinesExecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachinesExecBadRequest creates a MachinesExecBadRequest with default headers values
func NewMachinesExecBadRequest() *MachinesExecBadRequest {
	return &MachinesExecBadRequest{}
}

/*
MachinesExecBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MachinesExecBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this machines exec bad request response has a 2xx status code
func (o *MachinesExecBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this machines exec bad request response has a 3xx status code
func (o *MachinesExecBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machines exec bad request response has a 4xx status code
func (o *MachinesExecBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this machines exec bad request response has a 5xx status code
func (o *MachinesExecBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this machines exec bad request response a status code equal to that given
func (o *MachinesExecBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the machines exec bad request response
func (o *MachinesExecBadRequest) Code() int {
	return 400
}

func (o *MachinesExecBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}/exec][%d] machinesExecBadRequest  %+v", 400, o.Payload)
}

func (o *MachinesExecBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app_name}/machines/{machine_id}/exec][%d] machinesExecBadRequest  %+v", 400, o.Payload)
}

func (o *MachinesExecBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MachinesExecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
