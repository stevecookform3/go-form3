// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUsersUserIDCredentialsClientIDReader is a Reader for the DeleteUsersUserIDCredentialsClientID structure.
type DeleteUsersUserIDCredentialsClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersUserIDCredentialsClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUsersUserIDCredentialsClientIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsersUserIDCredentialsClientIDNoContent creates a DeleteUsersUserIDCredentialsClientIDNoContent with default headers values
func NewDeleteUsersUserIDCredentialsClientIDNoContent() *DeleteUsersUserIDCredentialsClientIDNoContent {
	return &DeleteUsersUserIDCredentialsClientIDNoContent{}
}

/*DeleteUsersUserIDCredentialsClientIDNoContent handles this case with default header values.

Credential deleted
*/
type DeleteUsersUserIDCredentialsClientIDNoContent struct {
}

func (o *DeleteUsersUserIDCredentialsClientIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials/{client_id}][%d] deleteUsersUserIdCredentialsClientIdNoContent ", 204)
}

func (o *DeleteUsersUserIDCredentialsClientIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
