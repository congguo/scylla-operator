// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindConfigBatchSizeFailThresholdInKbParams creates a new FindConfigBatchSizeFailThresholdInKbParams object
// with the default values initialized.
func NewFindConfigBatchSizeFailThresholdInKbParams() *FindConfigBatchSizeFailThresholdInKbParams {

	return &FindConfigBatchSizeFailThresholdInKbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigBatchSizeFailThresholdInKbParamsWithTimeout creates a new FindConfigBatchSizeFailThresholdInKbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigBatchSizeFailThresholdInKbParamsWithTimeout(timeout time.Duration) *FindConfigBatchSizeFailThresholdInKbParams {

	return &FindConfigBatchSizeFailThresholdInKbParams{

		timeout: timeout,
	}
}

// NewFindConfigBatchSizeFailThresholdInKbParamsWithContext creates a new FindConfigBatchSizeFailThresholdInKbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigBatchSizeFailThresholdInKbParamsWithContext(ctx context.Context) *FindConfigBatchSizeFailThresholdInKbParams {

	return &FindConfigBatchSizeFailThresholdInKbParams{

		Context: ctx,
	}
}

// NewFindConfigBatchSizeFailThresholdInKbParamsWithHTTPClient creates a new FindConfigBatchSizeFailThresholdInKbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigBatchSizeFailThresholdInKbParamsWithHTTPClient(client *http.Client) *FindConfigBatchSizeFailThresholdInKbParams {

	return &FindConfigBatchSizeFailThresholdInKbParams{
		HTTPClient: client,
	}
}

/*
FindConfigBatchSizeFailThresholdInKbParams contains all the parameters to send to the API endpoint
for the find config batch size fail threshold in kb operation typically these are written to a http.Request
*/
type FindConfigBatchSizeFailThresholdInKbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) WithTimeout(timeout time.Duration) *FindConfigBatchSizeFailThresholdInKbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) WithContext(ctx context.Context) *FindConfigBatchSizeFailThresholdInKbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) WithHTTPClient(client *http.Client) *FindConfigBatchSizeFailThresholdInKbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config batch size fail threshold in kb params
func (o *FindConfigBatchSizeFailThresholdInKbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigBatchSizeFailThresholdInKbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
