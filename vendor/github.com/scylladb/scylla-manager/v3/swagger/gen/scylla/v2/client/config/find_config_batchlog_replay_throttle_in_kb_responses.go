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

// FindConfigBatchlogReplayThrottleInKbReader is a Reader for the FindConfigBatchlogReplayThrottleInKb structure.
type FindConfigBatchlogReplayThrottleInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigBatchlogReplayThrottleInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigBatchlogReplayThrottleInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigBatchlogReplayThrottleInKbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigBatchlogReplayThrottleInKbOK creates a FindConfigBatchlogReplayThrottleInKbOK with default headers values
func NewFindConfigBatchlogReplayThrottleInKbOK() *FindConfigBatchlogReplayThrottleInKbOK {
	return &FindConfigBatchlogReplayThrottleInKbOK{}
}

/*
FindConfigBatchlogReplayThrottleInKbOK handles this case with default header values.

Config value
*/
type FindConfigBatchlogReplayThrottleInKbOK struct {
	Payload int64
}

func (o *FindConfigBatchlogReplayThrottleInKbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigBatchlogReplayThrottleInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigBatchlogReplayThrottleInKbDefault creates a FindConfigBatchlogReplayThrottleInKbDefault with default headers values
func NewFindConfigBatchlogReplayThrottleInKbDefault(code int) *FindConfigBatchlogReplayThrottleInKbDefault {
	return &FindConfigBatchlogReplayThrottleInKbDefault{
		_statusCode: code,
	}
}

/*
FindConfigBatchlogReplayThrottleInKbDefault handles this case with default header values.

unexpected error
*/
type FindConfigBatchlogReplayThrottleInKbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config batchlog replay throttle in kb default response
func (o *FindConfigBatchlogReplayThrottleInKbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigBatchlogReplayThrottleInKbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigBatchlogReplayThrottleInKbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigBatchlogReplayThrottleInKbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
