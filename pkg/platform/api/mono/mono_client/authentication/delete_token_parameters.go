// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewDeleteTokenParams creates a new DeleteTokenParams object
// with the default values initialized.
func NewDeleteTokenParams() *DeleteTokenParams {
	var ()
	return &DeleteTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTokenParamsWithTimeout creates a new DeleteTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTokenParamsWithTimeout(timeout time.Duration) *DeleteTokenParams {
	var ()
	return &DeleteTokenParams{

		timeout: timeout,
	}
}

// NewDeleteTokenParamsWithContext creates a new DeleteTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTokenParamsWithContext(ctx context.Context) *DeleteTokenParams {
	var ()
	return &DeleteTokenParams{

		Context: ctx,
	}
}

// NewDeleteTokenParamsWithHTTPClient creates a new DeleteTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTokenParamsWithHTTPClient(client *http.Client) *DeleteTokenParams {
	var ()
	return &DeleteTokenParams{
		HTTPClient: client,
	}
}

/*DeleteTokenParams contains all the parameters to send to the API endpoint
for the delete token operation typically these are written to a http.Request
*/
type DeleteTokenParams struct {

	/*TokenID
	  tokenID of desired API Token

	*/
	TokenID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete token params
func (o *DeleteTokenParams) WithTimeout(timeout time.Duration) *DeleteTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete token params
func (o *DeleteTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete token params
func (o *DeleteTokenParams) WithContext(ctx context.Context) *DeleteTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete token params
func (o *DeleteTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete token params
func (o *DeleteTokenParams) WithHTTPClient(client *http.Client) *DeleteTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete token params
func (o *DeleteTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTokenID adds the tokenID to the delete token params
func (o *DeleteTokenParams) WithTokenID(tokenID strfmt.UUID) *DeleteTokenParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the delete token params
func (o *DeleteTokenParams) SetTokenID(tokenID strfmt.UUID) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tokenID
	if err := r.SetPathParam("tokenID", o.TokenID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
