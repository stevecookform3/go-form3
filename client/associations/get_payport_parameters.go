// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewGetPayportParams creates a new GetPayportParams object
// with the default values initialized.
func NewGetPayportParams() *GetPayportParams {
	var ()
	return &GetPayportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPayportParamsWithTimeout creates a new GetPayportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPayportParamsWithTimeout(timeout time.Duration) *GetPayportParams {
	var ()
	return &GetPayportParams{

		timeout: timeout,
	}
}

// NewGetPayportParamsWithContext creates a new GetPayportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPayportParamsWithContext(ctx context.Context) *GetPayportParams {
	var ()
	return &GetPayportParams{

		Context: ctx,
	}
}

// NewGetPayportParamsWithHTTPClient creates a new GetPayportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPayportParamsWithHTTPClient(client *http.Client) *GetPayportParams {
	var ()
	return &GetPayportParams{
		HTTPClient: client,
	}
}

/*GetPayportParams contains all the parameters to send to the API endpoint
for the get payport operation typically these are written to a http.Request
*/
type GetPayportParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payport params
func (o *GetPayportParams) WithTimeout(timeout time.Duration) *GetPayportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payport params
func (o *GetPayportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payport params
func (o *GetPayportParams) WithContext(ctx context.Context) *GetPayportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payport params
func (o *GetPayportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payport params
func (o *GetPayportParams) WithHTTPClient(client *http.Client) *GetPayportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payport params
func (o *GetPayportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get payport params
func (o *GetPayportParams) WithFilterOrganisationID(filterOrganisationID *strfmt.UUID) *GetPayportParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get payport params
func (o *GetPayportParams) SetFilterOrganisationID(filterOrganisationID *strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPayportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterOrganisationID != nil {

		// query param filter[organisation_id]
		var qrFilterOrganisationID strfmt.UUID
		if o.FilterOrganisationID != nil {
			qrFilterOrganisationID = *o.FilterOrganisationID
		}
		qFilterOrganisationID := qrFilterOrganisationID.String()
		if qFilterOrganisationID != "" {
			if err := r.SetQueryParam("filter[organisation_id]", qFilterOrganisationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}