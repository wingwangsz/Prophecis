// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "mlss-controlcenter-go/pkg/models"
)

// NewUpdateNamespaceParams creates a new UpdateNamespaceParams object
// no default values defined in spec.
func NewUpdateNamespaceParams() UpdateNamespaceParams {

	return UpdateNamespaceParams{}
}

// UpdateNamespaceParams contains all the bound params for the update namespace operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateNamespace
type UpdateNamespaceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The namespace Request
	  Required: true
	  In: body
	*/
	Namespace *models.NamespaceRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateNamespaceParams() beforehand.
func (o *UpdateNamespaceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NamespaceRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("namespace", "body"))
			} else {
				res = append(res, errors.NewParseError("namespace", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Namespace = &body
			}
		}
	} else {
		res = append(res, errors.Required("namespace", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}