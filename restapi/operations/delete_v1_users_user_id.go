// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteV1UsersUserIDHandlerFunc turns a function with the right signature into a delete v1 users user ID handler
type DeleteV1UsersUserIDHandlerFunc func(DeleteV1UsersUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteV1UsersUserIDHandlerFunc) Handle(params DeleteV1UsersUserIDParams) middleware.Responder {
	return fn(params)
}

// DeleteV1UsersUserIDHandler interface for that can handle valid delete v1 users user ID params
type DeleteV1UsersUserIDHandler interface {
	Handle(DeleteV1UsersUserIDParams) middleware.Responder
}

// NewDeleteV1UsersUserID creates a new http.Handler for the delete v1 users user ID operation
func NewDeleteV1UsersUserID(ctx *middleware.Context, handler DeleteV1UsersUserIDHandler) *DeleteV1UsersUserID {
	return &DeleteV1UsersUserID{Context: ctx, Handler: handler}
}

/*DeleteV1UsersUserID swagger:route DELETE /v1/users/{userId} deleteV1UsersUserId

Removes the User data of given user id

Delete user data of given user id
```
DELETE /api/v1/users/{userId}
```


*/
type DeleteV1UsersUserID struct {
	Context *middleware.Context
	Handler DeleteV1UsersUserIDHandler
}

func (o *DeleteV1UsersUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteV1UsersUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
