// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1UsersUserIDParams creates a new GetV1UsersUserIDParams object
// no default values defined in spec.
func NewGetV1UsersUserIDParams() GetV1UsersUserIDParams {

	return GetV1UsersUserIDParams{}
}

// GetV1UsersUserIDParams contains all the bound params for the get v1 users user ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1UsersUserID
type GetV1UsersUserIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*User id
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1UsersUserIDParams() beforehand.
func (o *GetV1UsersUserIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("userId")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *GetV1UsersUserIDParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.UserID = raw

	return nil
}
