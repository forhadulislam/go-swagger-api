// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/forhadulislam/go-swagger-api/models"
)

// PostV1UsersOKCode is the HTTP code returned for type PostV1UsersOK
const PostV1UsersOKCode int = 200

/*PostV1UsersOK New user created successfully

swagger:response postV1UsersOK
*/
type PostV1UsersOK struct {

	/*
	  In: Body
	*/
	Payload models.UserID `json:"body,omitempty"`
}

// NewPostV1UsersOK creates PostV1UsersOK with default headers values
func NewPostV1UsersOK() *PostV1UsersOK {

	return &PostV1UsersOK{}
}

// WithPayload adds the payload to the post v1 users o k response
func (o *PostV1UsersOK) WithPayload(payload models.UserID) *PostV1UsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 users o k response
func (o *PostV1UsersOK) SetPayload(payload models.UserID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1UsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostV1UsersBadRequestCode is the HTTP code returned for type PostV1UsersBadRequest
const PostV1UsersBadRequestCode int = 400

/*PostV1UsersBadRequest Invalid request

swagger:response postV1UsersBadRequest
*/
type PostV1UsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewPostV1UsersBadRequest creates PostV1UsersBadRequest with default headers values
func NewPostV1UsersBadRequest() *PostV1UsersBadRequest {

	return &PostV1UsersBadRequest{}
}

// WithPayload adds the payload to the post v1 users bad request response
func (o *PostV1UsersBadRequest) WithPayload(payload *models.Status) *PostV1UsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 users bad request response
func (o *PostV1UsersBadRequest) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1UsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1UsersInternalServerErrorCode is the HTTP code returned for type PostV1UsersInternalServerError
const PostV1UsersInternalServerErrorCode int = 500

/*PostV1UsersInternalServerError Error while creating new user

swagger:response postV1UsersInternalServerError
*/
type PostV1UsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewPostV1UsersInternalServerError creates PostV1UsersInternalServerError with default headers values
func NewPostV1UsersInternalServerError() *PostV1UsersInternalServerError {

	return &PostV1UsersInternalServerError{}
}

// WithPayload adds the payload to the post v1 users internal server error response
func (o *PostV1UsersInternalServerError) WithPayload(payload *models.Status) *PostV1UsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 users internal server error response
func (o *PostV1UsersInternalServerError) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1UsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
