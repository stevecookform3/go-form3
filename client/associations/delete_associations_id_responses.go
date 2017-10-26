// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAssociationsIDReader is a Reader for the DeleteAssociationsID structure.
type DeleteAssociationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAssociationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAssociationsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAssociationsIDNoContent creates a DeleteAssociationsIDNoContent with default headers values
func NewDeleteAssociationsIDNoContent() *DeleteAssociationsIDNoContent {
	return &DeleteAssociationsIDNoContent{}
}

/*DeleteAssociationsIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteAssociationsIDNoContent struct {
}

func (o *DeleteAssociationsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /associations/{id}][%d] deleteAssociationsIdNoContent ", 204)
}

func (o *DeleteAssociationsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
