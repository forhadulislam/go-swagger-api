// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/forhadulislam/go-swagger-api/models"
)

// NewPostV1UsersParams creates a new PostV1UsersParams object
// no default values defined in spec.
func NewPostV1UsersParams() PostV1UsersParams {

	return PostV1UsersParams{}
}

// PostV1UsersParams contains all the bound params for the post v1 users operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostV1Users
type PostV1UsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Parameters of users information
	  Required: true
	  In: body
	*/
	UserRequest *models.UserRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostV1UsersParams() beforehand.
func (o *PostV1UsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("userRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UserRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("userRequest", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
