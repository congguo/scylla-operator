// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceViewBuildStatusesByKeyspaceAndViewGetReader is a Reader for the StorageServiceViewBuildStatusesByKeyspaceAndViewGet structure.
type StorageServiceViewBuildStatusesByKeyspaceAndViewGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetOK creates a StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK with default headers values
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetOK() *StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK {
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK{}
}

/*
StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK handles this case with default header values.

Success
*/
type StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault creates a StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault with default headers values
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault(code int) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault {
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service view build statuses by keyspace and view get default response
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
