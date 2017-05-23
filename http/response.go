package http

import (
	"github.com/pkg/errors"
	"net/http"
)

// abstractions for an http response
type ResponseSource interface {
	// status for this response
	Status() int

	// get a header from this response
	//
	// implementations should return zero string if
	// no header of this name is in the response.
	Header(name string) string

	// get the body of this response
	//
	// implementations should return zero-length byte array
	// if the response does not contain a body.
	Body() []byte

	// write all contents to the actual `http.ResponseWriter`
	Write(w http.ResponseWriter) error
}

// default implementation of `ResponseSource`
type HttpResponse struct {
	StatusCode int
	Headers    map[string]string
	Data       []byte
}

func (resp *HttpResponse) Status() int {
	return resp.StatusCode
}

func (resp *HttpResponse) Header(name string) string {
	return resp.Headers[name]
}

func (resp *HttpResponse) Body() []byte {
	return resp.Data
}

func (resp *HttpResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(resp.Status())
	for k, v := range resp.Headers {
		if len(k) > 0 && len(v) > 0 {
			w.Header().Set(k, v)
		}
	}
	if _, err := w.Write(resp.Data); err != nil {
		return errors.WithStack(err)
	} else {
		return nil
	}
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{
		Headers: make(map[string]string),
		Data:    make([]byte, 0),
	}
}
