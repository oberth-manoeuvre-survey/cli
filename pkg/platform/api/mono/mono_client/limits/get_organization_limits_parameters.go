// Code generated by go-swagger; DO NOT EDIT.

package limits

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

// NewGetOrganizationLimitsParams creates a new GetOrganizationLimitsParams object
// with the default values initialized.
func NewGetOrganizationLimitsParams() *GetOrganizationLimitsParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &GetOrganizationLimitsParams{
		IdentifierType: &identifierTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationLimitsParamsWithTimeout creates a new GetOrganizationLimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationLimitsParamsWithTimeout(timeout time.Duration) *GetOrganizationLimitsParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &GetOrganizationLimitsParams{
		IdentifierType: &identifierTypeDefault,

		timeout: timeout,
	}
}

// NewGetOrganizationLimitsParamsWithContext creates a new GetOrganizationLimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationLimitsParamsWithContext(ctx context.Context) *GetOrganizationLimitsParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &GetOrganizationLimitsParams{
		IdentifierType: &identifierTypeDefault,

		Context: ctx,
	}
}

// NewGetOrganizationLimitsParamsWithHTTPClient creates a new GetOrganizationLimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationLimitsParamsWithHTTPClient(client *http.Client) *GetOrganizationLimitsParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &GetOrganizationLimitsParams{
		IdentifierType: &identifierTypeDefault,
		HTTPClient:     client,
	}
}

/*GetOrganizationLimitsParams contains all the parameters to send to the API endpoint
for the get organization limits operation typically these are written to a http.Request
*/
type GetOrganizationLimitsParams struct {

	/*IdentifierType
	  what kind of thing the provided organizationIdentifier is

	*/
	IdentifierType *string
	/*OrganizationIdentifier
	  identifier (URLname, by default) of the desired organization

	*/
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization limits params
func (o *GetOrganizationLimitsParams) WithTimeout(timeout time.Duration) *GetOrganizationLimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization limits params
func (o *GetOrganizationLimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization limits params
func (o *GetOrganizationLimitsParams) WithContext(ctx context.Context) *GetOrganizationLimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization limits params
func (o *GetOrganizationLimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization limits params
func (o *GetOrganizationLimitsParams) WithHTTPClient(client *http.Client) *GetOrganizationLimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization limits params
func (o *GetOrganizationLimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifierType adds the identifierType to the get organization limits params
func (o *GetOrganizationLimitsParams) WithIdentifierType(identifierType *string) *GetOrganizationLimitsParams {
	o.SetIdentifierType(identifierType)
	return o
}

// SetIdentifierType adds the identifierType to the get organization limits params
func (o *GetOrganizationLimitsParams) SetIdentifierType(identifierType *string) {
	o.IdentifierType = identifierType
}

// WithOrganizationIdentifier adds the organizationIdentifier to the get organization limits params
func (o *GetOrganizationLimitsParams) WithOrganizationIdentifier(organizationIdentifier string) *GetOrganizationLimitsParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the get organization limits params
func (o *GetOrganizationLimitsParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationLimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IdentifierType != nil {

		// query param identifierType
		var qrIdentifierType string
		if o.IdentifierType != nil {
			qrIdentifierType = *o.IdentifierType
		}
		qIdentifierType := qrIdentifierType
		if qIdentifierType != "" {
			if err := r.SetQueryParam("identifierType", qIdentifierType); err != nil {
				return err
			}
		}

	}

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
