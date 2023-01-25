// Code generated by goa v3.11.0, DO NOT EDIT.
//
// category HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	category "github.com/tektoncd/hub/api/gen/category"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "category" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data []*CategoryResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ListInternalErrorResponseBody is the type of the "category" service "list"
// endpoint HTTP response body for the "internal-error" error.
type ListInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	// ID is the unique id of the category
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of category
	Name string `form:"name" json:"name" xml:"name"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "category" service.
func NewListResponseBody(res *category.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Data != nil {
		body.Data = make([]*CategoryResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalCategoryCategoryToCategoryResponseBody(val)
		}
	}
	return body
}

// NewListInternalErrorResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "category" service.
func NewListInternalErrorResponseBody(res *goa.ServiceError) *ListInternalErrorResponseBody {
	body := &ListInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
