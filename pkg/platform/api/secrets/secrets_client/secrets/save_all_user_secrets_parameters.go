// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

	secrets_models "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
)

// NewSaveAllUserSecretsParams creates a new SaveAllUserSecretsParams object
// with the default values initialized.
func NewSaveAllUserSecretsParams() *SaveAllUserSecretsParams {
	var ()
	return &SaveAllUserSecretsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveAllUserSecretsParamsWithTimeout creates a new SaveAllUserSecretsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveAllUserSecretsParamsWithTimeout(timeout time.Duration) *SaveAllUserSecretsParams {
	var ()
	return &SaveAllUserSecretsParams{

		timeout: timeout,
	}
}

// NewSaveAllUserSecretsParamsWithContext creates a new SaveAllUserSecretsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveAllUserSecretsParamsWithContext(ctx context.Context) *SaveAllUserSecretsParams {
	var ()
	return &SaveAllUserSecretsParams{

		Context: ctx,
	}
}

// NewSaveAllUserSecretsParamsWithHTTPClient creates a new SaveAllUserSecretsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveAllUserSecretsParamsWithHTTPClient(client *http.Client) *SaveAllUserSecretsParams {
	var ()
	return &SaveAllUserSecretsParams{
		HTTPClient: client,
	}
}

/*SaveAllUserSecretsParams contains all the parameters to send to the API endpoint
for the save all user secrets operation typically these are written to a http.Request
*/
type SaveAllUserSecretsParams struct {

	/*OrganizationID
	  Organization-id of organization to store the secrets against

	*/
	OrganizationID strfmt.UUID
	/*UserSecrets
	  Collection of secrets to create or update

	*/
	UserSecrets []*secrets_models.UserSecretChange

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save all user secrets params
func (o *SaveAllUserSecretsParams) WithTimeout(timeout time.Duration) *SaveAllUserSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save all user secrets params
func (o *SaveAllUserSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save all user secrets params
func (o *SaveAllUserSecretsParams) WithContext(ctx context.Context) *SaveAllUserSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save all user secrets params
func (o *SaveAllUserSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save all user secrets params
func (o *SaveAllUserSecretsParams) WithHTTPClient(client *http.Client) *SaveAllUserSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save all user secrets params
func (o *SaveAllUserSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the save all user secrets params
func (o *SaveAllUserSecretsParams) WithOrganizationID(organizationID strfmt.UUID) *SaveAllUserSecretsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the save all user secrets params
func (o *SaveAllUserSecretsParams) SetOrganizationID(organizationID strfmt.UUID) {
	o.OrganizationID = organizationID
}

// WithUserSecrets adds the userSecrets to the save all user secrets params
func (o *SaveAllUserSecretsParams) WithUserSecrets(userSecrets []*secrets_models.UserSecretChange) *SaveAllUserSecretsParams {
	o.SetUserSecrets(userSecrets)
	return o
}

// SetUserSecrets adds the userSecrets to the save all user secrets params
func (o *SaveAllUserSecretsParams) SetUserSecrets(userSecrets []*secrets_models.UserSecretChange) {
	o.UserSecrets = userSecrets
}

// WriteToRequest writes these params to a swagger request
func (o *SaveAllUserSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationID
	if err := r.SetPathParam("organizationID", o.OrganizationID.String()); err != nil {
		return err
	}

	if o.UserSecrets != nil {
		if err := r.SetBodyParam(o.UserSecrets); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
