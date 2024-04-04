// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/agent/models"
)

// CoreStatsResetReader is a Reader for the CoreStatsReset structure.
type CoreStatsResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreStatsResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreStatsResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoreStatsResetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoreStatsResetOK creates a CoreStatsResetOK with default headers values
func NewCoreStatsResetOK() *CoreStatsResetOK {
	return &CoreStatsResetOK{}
}

/*
CoreStatsResetOK handles this case with default header values.

Empty object
*/
type CoreStatsResetOK struct {
	Payload interface{}
	JobID   int64
}

func (o *CoreStatsResetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CoreStatsResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewCoreStatsResetDefault creates a CoreStatsResetDefault with default headers values
func NewCoreStatsResetDefault(code int) *CoreStatsResetDefault {
	return &CoreStatsResetDefault{
		_statusCode: code,
	}
}

/*
CoreStatsResetDefault handles this case with default header values.

Server error
*/
type CoreStatsResetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the core stats reset default response
func (o *CoreStatsResetDefault) Code() int {
	return o._statusCode
}

func (o *CoreStatsResetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoreStatsResetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *CoreStatsResetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
