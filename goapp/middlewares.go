package goapp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// exmaple of custom httprouter middleware
func httprouterSample(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h(w, r, ps)
	}
}

// example of custom net/http middleware
func sample(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
