// Code generated by go-swagger; DO NOT EDIT.

package queries

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
	"github.com/go-openapi/swag"
)

// NewGetQueriesParams creates a new GetQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueriesParams() *GetQueriesParams {
	return &GetQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueriesParamsWithTimeout creates a new GetQueriesParams object
// with the ability to set a timeout on a request.
func NewGetQueriesParamsWithTimeout(timeout time.Duration) *GetQueriesParams {
	return &GetQueriesParams{
		timeout: timeout,
	}
}

// NewGetQueriesParamsWithContext creates a new GetQueriesParams object
// with the ability to set a context for a request.
func NewGetQueriesParamsWithContext(ctx context.Context) *GetQueriesParams {
	return &GetQueriesParams{
		Context: ctx,
	}
}

// NewGetQueriesParamsWithHTTPClient creates a new GetQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueriesParamsWithHTTPClient(client *http.Client) *GetQueriesParams {
	return &GetQueriesParams{
		HTTPClient: client,
	}
}

/* GetQueriesParams contains all the parameters to send to the API endpoint
   for the get queries operation.

   Typically these are written to a http.Request.
*/
type GetQueriesParams struct {

	/* PageSize.

	   Numeric ID to limit number of results

	   Default: 100
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueriesParams) WithDefaults() *GetQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueriesParams) SetDefaults() {
	var (
		pageSizeDefault = int64(100)
	)

	val := GetQueriesParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get queries params
func (o *GetQueriesParams) WithTimeout(timeout time.Duration) *GetQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queries params
func (o *GetQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queries params
func (o *GetQueriesParams) WithContext(ctx context.Context) *GetQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queries params
func (o *GetQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queries params
func (o *GetQueriesParams) WithHTTPClient(client *http.Client) *GetQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queries params
func (o *GetQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the get queries params
func (o *GetQueriesParams) WithPageSize(pageSize *int64) *GetQueriesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get queries params
func (o *GetQueriesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
