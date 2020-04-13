// Code generated by go-swagger; DO NOT EDIT.

package sessions

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

// NewGetSessionParams creates a new GetSessionParams object
// with the default values initialized.
func NewGetSessionParams() *GetSessionParams {
	var ()
	return &GetSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSessionParamsWithTimeout creates a new GetSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSessionParamsWithTimeout(timeout time.Duration) *GetSessionParams {
	var ()
	return &GetSessionParams{

		timeout: timeout,
	}
}

// NewGetSessionParamsWithContext creates a new GetSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSessionParamsWithContext(ctx context.Context) *GetSessionParams {
	var ()
	return &GetSessionParams{

		Context: ctx,
	}
}

// NewGetSessionParamsWithHTTPClient creates a new GetSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSessionParamsWithHTTPClient(client *http.Client) *GetSessionParams {
	var ()
	return &GetSessionParams{
		HTTPClient: client,
	}
}

/*GetSessionParams contains all the parameters to send to the API endpoint
for the get session operation typically these are written to a http.Request
*/
type GetSessionParams struct {

	/*SessionID
	  Unique ID of session

	*/
	SessionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get session params
func (o *GetSessionParams) WithTimeout(timeout time.Duration) *GetSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get session params
func (o *GetSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get session params
func (o *GetSessionParams) WithContext(ctx context.Context) *GetSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get session params
func (o *GetSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get session params
func (o *GetSessionParams) WithHTTPClient(client *http.Client) *GetSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get session params
func (o *GetSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionID adds the sessionID to the get session params
func (o *GetSessionParams) WithSessionID(sessionID strfmt.UUID) *GetSessionParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get session params
func (o *GetSessionParams) SetSessionID(sessionID strfmt.UUID) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionID
	if err := r.SetPathParam("sessionID", o.SessionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
