// Package middleware includes handlers taht do some type of preprocessing beofr the main handlers.
package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger logs the current server request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

// Authenticate attempts to authenticate the user if they have the appropriate cookie.
func Authenticate() {
}
