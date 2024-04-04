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

// NewFindConfigReplaceNodeParams creates a new FindConfigReplaceNodeParams object
// with the default values initialized.
func NewFindConfigReplaceNodeParams() *FindConfigReplaceNodeParams {

	return &FindConfigReplaceNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigReplaceNodeParamsWithTimeout creates a new FindConfigReplaceNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigReplaceNodeParamsWithTimeout(timeout time.Duration) *FindConfigReplaceNodeParams {

	return &FindConfigReplaceNodeParams{

		timeout: timeout,
	}
}

// NewFindConfigReplaceNodeParamsWithContext creates a new FindConfigReplaceNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigReplaceNodeParamsWithContext(ctx context.Context) *FindConfigReplaceNodeParams {

	return &FindConfigReplaceNodeParams{

		Context: ctx,
	}
}

// NewFindConfigReplaceNodeParamsWithHTTPClient creates a new FindConfigReplaceNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigReplaceNodeParamsWithHTTPClient(client *http.Client) *FindConfigReplaceNodeParams {

	return &FindConfigReplaceNodeParams{
		HTTPClient: client,
	}
}

/*
FindConfigReplaceNodeParams contains all the parameters to send to the API endpoint
for the find config replace node operation typically these are written to a http.Request
*/
type FindConfigReplaceNodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config replace node params
func (o *FindConfigReplaceNodeParams) WithTimeout(timeout time.Duration) *FindConfigReplaceNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config replace node params
func (o *FindConfigReplaceNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config replace node params
func (o *FindConfigReplaceNodeParams) WithContext(ctx context.Context) *FindConfigReplaceNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config replace node params
func (o *FindConfigReplaceNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config replace node params
func (o *FindConfigReplaceNodeParams) WithHTTPClient(client *http.Client) *FindConfigReplaceNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config replace node params
func (o *FindConfigReplaceNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigReplaceNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
