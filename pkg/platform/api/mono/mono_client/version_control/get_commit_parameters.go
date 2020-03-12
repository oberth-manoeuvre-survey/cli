// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCommitParams creates a new GetCommitParams object
// with the default values initialized.
func NewGetCommitParams() *GetCommitParams {
	var ()
	return &GetCommitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommitParamsWithTimeout creates a new GetCommitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCommitParamsWithTimeout(timeout time.Duration) *GetCommitParams {
	var ()
	return &GetCommitParams{

		timeout: timeout,
	}
}

// NewGetCommitParamsWithContext creates a new GetCommitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCommitParamsWithContext(ctx context.Context) *GetCommitParams {
	var ()
	return &GetCommitParams{

		Context: ctx,
	}
}

// NewGetCommitParamsWithHTTPClient creates a new GetCommitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCommitParamsWithHTTPClient(client *http.Client) *GetCommitParams {
	var ()
	return &GetCommitParams{
		HTTPClient: client,
	}
}

/*GetCommitParams contains all the parameters to send to the API endpoint
for the get commit operation typically these are written to a http.Request
*/
type GetCommitParams struct {

	/*CommitID*/
	CommitID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get commit params
func (o *GetCommitParams) WithTimeout(timeout time.Duration) *GetCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get commit params
func (o *GetCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get commit params
func (o *GetCommitParams) WithContext(ctx context.Context) *GetCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get commit params
func (o *GetCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get commit params
func (o *GetCommitParams) WithHTTPClient(client *http.Client) *GetCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get commit params
func (o *GetCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommitID adds the commitID to the get commit params
func (o *GetCommitParams) WithCommitID(commitID strfmt.UUID) *GetCommitParams {
	o.SetCommitID(commitID)
	return o
}

// SetCommitID adds the commitId to the get commit params
func (o *GetCommitParams) SetCommitID(commitID strfmt.UUID) {
	o.CommitID = commitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param commitID
	if err := r.SetPathParam("commitID", o.CommitID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
