package http

import "net/http"

// An endpoint handler.
//
// It takes in a request source and returns a response source.
// It is recommended to panic any errors and have downstream interceptors deal
// with them.
type RequestHandler func(req RequestSource) ResponseSource

// Creates a standard `http.HandlerFunc`.
//
// Adapts RequestHandler to the standard Go API.
func CreateHandler(handler RequestHandler, bindParamFunc ExtractBindParamFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := NewDefaultRequestSource(r, bindParamFunc)
		resp := handler(req)
		resp.Write(w)
	})
}

// An endpoint filter
//
// It has the choice to perform some operations either before or after invoking
// the endpoint handler.
type EndpointFilter func(next RequestHandler) ResponseSource

// Adds panic recovery capabilities to handler
func AddPanicRecovery(handler RequestHandler, recoverFunc func(r interface{}) ResponseSource) RequestHandler {
	return func(req RequestSource) (resp ResponseSource) {
		defer func() {
			if r := recover(); r != nil {
				resp = recoverFunc(r)
			}
		}()
		return handler(req)
	}
}
