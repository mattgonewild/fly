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

// MachinesListReader is a Reader for the MachinesList structure.
type MachinesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachinesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachinesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app_name}/machines] Machines_list", response, response.Code())
	}
}

// NewMachinesListOK creates a MachinesListOK with default headers values
func NewMachinesListOK() *MachinesListOK {
	return &MachinesListOK{}
}

/*
MachinesListOK describes a response with status code 200, with default header values.

OK
*/
type MachinesListOK struct {
	Payload []*models.Machine
}

// IsSuccess returns true when this machines list o k response has a 2xx status code
func (o *MachinesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machines list o k response has a 3xx status code
func (o *MachinesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machines list o k response has a 4xx status code
func (o *MachinesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machines list o k response has a 5xx status code
func (o *MachinesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machines list o k response a status code equal to that given
func (o *MachinesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machines list o k response
func (o *MachinesListOK) Code() int {
	return 200
}

func (o *MachinesListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_name}/machines][%d] machinesListOK  %+v", 200, o.Payload)
}

func (o *MachinesListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app_name}/machines][%d] machinesListOK  %+v", 200, o.Payload)
}

func (o *MachinesListOK) GetPayload() []*models.Machine {
	return o.Payload
}

func (o *MachinesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
