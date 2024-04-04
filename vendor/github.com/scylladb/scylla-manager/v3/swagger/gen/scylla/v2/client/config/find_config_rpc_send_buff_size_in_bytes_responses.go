// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v2/models"
)

// FindConfigRPCSendBuffSizeInBytesReader is a Reader for the FindConfigRPCSendBuffSizeInBytes structure.
type FindConfigRPCSendBuffSizeInBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRPCSendBuffSizeInBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRPCSendBuffSizeInBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRPCSendBuffSizeInBytesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRPCSendBuffSizeInBytesOK creates a FindConfigRPCSendBuffSizeInBytesOK with default headers values
func NewFindConfigRPCSendBuffSizeInBytesOK() *FindConfigRPCSendBuffSizeInBytesOK {
	return &FindConfigRPCSendBuffSizeInBytesOK{}
}

/*
FindConfigRPCSendBuffSizeInBytesOK handles this case with default header values.

Config value
*/
type FindConfigRPCSendBuffSizeInBytesOK struct {
	Payload int64
}

func (o *FindConfigRPCSendBuffSizeInBytesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigRPCSendBuffSizeInBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRPCSendBuffSizeInBytesDefault creates a FindConfigRPCSendBuffSizeInBytesDefault with default headers values
func NewFindConfigRPCSendBuffSizeInBytesDefault(code int) *FindConfigRPCSendBuffSizeInBytesDefault {
	return &FindConfigRPCSendBuffSizeInBytesDefault{
		_statusCode: code,
	}
}

/*
FindConfigRPCSendBuffSizeInBytesDefault handles this case with default header values.

unexpected error
*/
type FindConfigRPCSendBuffSizeInBytesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config rpc send buff size in bytes default response
func (o *FindConfigRPCSendBuffSizeInBytesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRPCSendBuffSizeInBytesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRPCSendBuffSizeInBytesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRPCSendBuffSizeInBytesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
