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

// VolumesListSnapshotsReader is a Reader for the VolumesListSnapshots structure.
type VolumesListSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumesListSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumesListSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app_name}/volumes/{volume_id}/snapshots] Volumes_list_snapshots", response, response.Code())
	}
}

// NewVolumesListSnapshotsOK creates a VolumesListSnapshotsOK with default headers values
func NewVolumesListSnapshotsOK() *VolumesListSnapshotsOK {
	return &VolumesListSnapshotsOK{}
}

/*
VolumesListSnapshotsOK describes a response with status code 200, with default header values.

OK
*/
type VolumesListSnapshotsOK struct {
	Payload []*models.VolumeSnapshot
}

// IsSuccess returns true when this volumes list snapshots o k response has a 2xx status code
func (o *VolumesListSnapshotsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volumes list snapshots o k response has a 3xx status code
func (o *VolumesListSnapshotsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volumes list snapshots o k response has a 4xx status code
func (o *VolumesListSnapshotsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volumes list snapshots o k response has a 5xx status code
func (o *VolumesListSnapshotsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volumes list snapshots o k response a status code equal to that given
func (o *VolumesListSnapshotsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volumes list snapshots o k response
func (o *VolumesListSnapshotsOK) Code() int {
	return 200
}

func (o *VolumesListSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_name}/volumes/{volume_id}/snapshots][%d] volumesListSnapshotsOK  %+v", 200, o.Payload)
}

func (o *VolumesListSnapshotsOK) String() string {
	return fmt.Sprintf("[GET /apps/{app_name}/volumes/{volume_id}/snapshots][%d] volumesListSnapshotsOK  %+v", 200, o.Payload)
}

func (o *VolumesListSnapshotsOK) GetPayload() []*models.VolumeSnapshot {
	return o.Payload
}

func (o *VolumesListSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
