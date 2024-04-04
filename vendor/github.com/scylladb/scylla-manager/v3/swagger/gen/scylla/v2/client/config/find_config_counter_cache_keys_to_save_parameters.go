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

// NewFindConfigCounterCacheKeysToSaveParams creates a new FindConfigCounterCacheKeysToSaveParams object
// with the default values initialized.
func NewFindConfigCounterCacheKeysToSaveParams() *FindConfigCounterCacheKeysToSaveParams {

	return &FindConfigCounterCacheKeysToSaveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCounterCacheKeysToSaveParamsWithTimeout creates a new FindConfigCounterCacheKeysToSaveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCounterCacheKeysToSaveParamsWithTimeout(timeout time.Duration) *FindConfigCounterCacheKeysToSaveParams {

	return &FindConfigCounterCacheKeysToSaveParams{

		timeout: timeout,
	}
}

// NewFindConfigCounterCacheKeysToSaveParamsWithContext creates a new FindConfigCounterCacheKeysToSaveParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCounterCacheKeysToSaveParamsWithContext(ctx context.Context) *FindConfigCounterCacheKeysToSaveParams {

	return &FindConfigCounterCacheKeysToSaveParams{

		Context: ctx,
	}
}

// NewFindConfigCounterCacheKeysToSaveParamsWithHTTPClient creates a new FindConfigCounterCacheKeysToSaveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCounterCacheKeysToSaveParamsWithHTTPClient(client *http.Client) *FindConfigCounterCacheKeysToSaveParams {

	return &FindConfigCounterCacheKeysToSaveParams{
		HTTPClient: client,
	}
}

/*
FindConfigCounterCacheKeysToSaveParams contains all the parameters to send to the API endpoint
for the find config counter cache keys to save operation typically these are written to a http.Request
*/
type FindConfigCounterCacheKeysToSaveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) WithTimeout(timeout time.Duration) *FindConfigCounterCacheKeysToSaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) WithContext(ctx context.Context) *FindConfigCounterCacheKeysToSaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) WithHTTPClient(client *http.Client) *FindConfigCounterCacheKeysToSaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config counter cache keys to save params
func (o *FindConfigCounterCacheKeysToSaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCounterCacheKeysToSaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
