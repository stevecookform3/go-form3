// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAccountconfigurationsIDReader is a Reader for the DeleteAccountconfigurationsID structure.
type DeleteAccountconfigurationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountconfigurationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAccountconfigurationsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountconfigurationsIDNoContent creates a DeleteAccountconfigurationsIDNoContent with default headers values
func NewDeleteAccountconfigurationsIDNoContent() *DeleteAccountconfigurationsIDNoContent {
	return &DeleteAccountconfigurationsIDNoContent{}
}

/*DeleteAccountconfigurationsIDNoContent handles this case with default header values.

AccountConfiguration deleted
*/
type DeleteAccountconfigurationsIDNoContent struct {
}

func (o *DeleteAccountconfigurationsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdNoContent ", 204)
}

func (o *DeleteAccountconfigurationsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
