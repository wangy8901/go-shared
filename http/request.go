package http

import (
	"net/http"
	"io/ioutil"
	"github.com/pkg/errors"
	"context"
)

// abstractions of an http request
type RequestSource interface {
	// request uri
	Uri() string

	// http method
	//
	// implementations must return standard
	// http verbs in upper case form.
	Method() string

	// retrieve http header by name
	//
	// name is case sensitive http header.
	// implementations must return zero string
	// for no existent header.
	Header(name string) string

	// retrieve url parameters by name
	//
	// name is the case sensitive url parameter name or
	// the bound url parameter placeholder name.
	// implementations must return zero string
	// for no existent parameter.
	Param(name string) string

	// body of the request
	//
	// implementations must return error and zero-length byte
	// array in case something went wrong during extracting
	// the body content.
	Body() ([]byte, error)

	// get the context for this request
	Context() context.Context

	// get the underlying request object (if feasible) for further operation
	Request() interface{}
}

// A function to extract bind parameters from url
//
// This puts the logic and knowledge of bind parameters name
// down to the implementations. Implementations must return
// zero string if there's no bind parameter by the name.
type ExtractBindParamFunc func(r *http.Request, name string) string

var (
	_ RequestSource = (*defaultRequestSource)(nil)
)

// implementation of `RequestSource` that wraps inside
// a `*http.Request`
type defaultRequestSource struct {
	req		*http.Request
	extractFunc	ExtractBindParamFunc
}

func (rs *defaultRequestSource) Uri() string {
	return rs.req.RequestURI
}

func (rs *defaultRequestSource) Method() string {
	return rs.req.Method
}

func (rs *defaultRequestSource) Header(name string) string {
	return rs.req.Header.Get(name)
}

func (rs *defaultRequestSource) Param(name string) string {
	if v := rs.req.URL.Query().Get(name); len(v) > 0 {
		return v
	} else if v := rs.extractFunc(rs.req, name); len(v) > 0 {
		return v
	} else {
		return ""
	}
}

func (rs *defaultRequestSource) Body() ([]byte, error) {
	if bytes, err := ioutil.ReadAll(rs.req.Body); err != nil {
		return nil, errors.WithStack(err)
	} else {
		return bytes, nil
	}
}

func (rs *defaultRequestSource) Context() context.Context {
	return rs.req.Context()
}

func (rs *defaultRequestSource) Request() interface{} {
	return rs.req
}

func NewDefaultRequestSource(req *http.Request, bindParamFunc ExtractBindParamFunc) RequestSource {
	return &defaultRequestSource{req:req, extractFunc:bindParamFunc}
}