// Code generated by go-swagger; DO NOT EDIT.

package limits

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

	"github.com/ewilde/go-form3/models"
)

// NewPatchLimitsIDParams creates a new PatchLimitsIDParams object
// with the default values initialized.
func NewPatchLimitsIDParams() *PatchLimitsIDParams {
	var ()
	return &PatchLimitsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLimitsIDParamsWithTimeout creates a new PatchLimitsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLimitsIDParamsWithTimeout(timeout time.Duration) *PatchLimitsIDParams {
	var ()
	return &PatchLimitsIDParams{

		timeout: timeout,
	}
}

// NewPatchLimitsIDParamsWithContext creates a new PatchLimitsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLimitsIDParamsWithContext(ctx context.Context) *PatchLimitsIDParams {
	var ()
	return &PatchLimitsIDParams{

		Context: ctx,
	}
}

// NewPatchLimitsIDParamsWithHTTPClient creates a new PatchLimitsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLimitsIDParamsWithHTTPClient(client *http.Client) *PatchLimitsIDParams {
	var ()
	return &PatchLimitsIDParams{
		HTTPClient: client,
	}
}

/*PatchLimitsIDParams contains all the parameters to send to the API endpoint
for the patch limits ID operation typically these are written to a http.Request
*/
type PatchLimitsIDParams struct {

	/*LimitAmendRequest*/
	LimitAmendRequest *models.LimitAmendment
	/*ID
	  Limit Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch limits ID params
func (o *PatchLimitsIDParams) WithTimeout(timeout time.Duration) *PatchLimitsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch limits ID params
func (o *PatchLimitsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch limits ID params
func (o *PatchLimitsIDParams) WithContext(ctx context.Context) *PatchLimitsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch limits ID params
func (o *PatchLimitsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch limits ID params
func (o *PatchLimitsIDParams) WithHTTPClient(client *http.Client) *PatchLimitsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch limits ID params
func (o *PatchLimitsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimitAmendRequest adds the limitAmendRequest to the patch limits ID params
func (o *PatchLimitsIDParams) WithLimitAmendRequest(limitAmendRequest *models.LimitAmendment) *PatchLimitsIDParams {
	o.SetLimitAmendRequest(limitAmendRequest)
	return o
}

// SetLimitAmendRequest adds the limitAmendRequest to the patch limits ID params
func (o *PatchLimitsIDParams) SetLimitAmendRequest(limitAmendRequest *models.LimitAmendment) {
	o.LimitAmendRequest = limitAmendRequest
}

// WithID adds the id to the patch limits ID params
func (o *PatchLimitsIDParams) WithID(id strfmt.UUID) *PatchLimitsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch limits ID params
func (o *PatchLimitsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLimitsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LimitAmendRequest != nil {
		if err := r.SetBodyParam(o.LimitAmendRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
