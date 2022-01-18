// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/recolabs/reco/redash-client/gen/models"
)

// GetQueriesIDReader is a Reader for the GetQueriesID structure.
type GetQueriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetQueriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueriesIDOK creates a GetQueriesIDOK with default headers values
func NewGetQueriesIDOK() *GetQueriesIDOK {
	return &GetQueriesIDOK{}
}

/* GetQueriesIDOK describes a response with status code 200, with default header values.

Returned query object
*/
type GetQueriesIDOK struct {
	Payload *models.Query
}

func (o *GetQueriesIDOK) Error() string {
	return fmt.Sprintf("[GET /queries/{id}][%d] getQueriesIdOK  %+v", 200, o.Payload)
}
func (o *GetQueriesIDOK) GetPayload() *models.Query {
	return o.Payload
}

func (o *GetQueriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Query)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueriesIDDefault creates a GetQueriesIDDefault with default headers values
func NewGetQueriesIDDefault(code int) *GetQueriesIDDefault {
	return &GetQueriesIDDefault{
		_statusCode: code,
	}
}

/* GetQueriesIDDefault describes a response with status code -1, with default header values.

error
*/
type GetQueriesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get queries ID default response
func (o *GetQueriesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetQueriesIDDefault) Error() string {
	return fmt.Sprintf("[GET /queries/{id}][%d] GetQueriesID default  %+v", o._statusCode, o.Payload)
}
func (o *GetQueriesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetQueriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
