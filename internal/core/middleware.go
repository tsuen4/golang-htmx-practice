package core

import (
	"log"
	"net/http"
)

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("%s %s\n", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)
	})
}
