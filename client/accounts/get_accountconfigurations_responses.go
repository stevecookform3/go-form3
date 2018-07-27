// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ewilde/go-form3/models"
)

// GetAccountconfigurationsReader is a Reader for the GetAccountconfigurations structure.
type GetAccountconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountconfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountconfigurationsOK creates a GetAccountconfigurationsOK with default headers values
func NewGetAccountconfigurationsOK() *GetAccountconfigurationsOK {
	return &GetAccountconfigurationsOK{}
}

/*GetAccountconfigurationsOK handles this case with default header values.

List of configuration details
*/
type GetAccountconfigurationsOK struct {
	Payload *models.AccountConfigurationDetailsListResponse
}

func (o *GetAccountconfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /accountconfigurations][%d] getAccountconfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetAccountconfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountConfigurationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
