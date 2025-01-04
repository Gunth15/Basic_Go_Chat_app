// Package middleware includes handlers taht do some type of preprocessing beofr the main handlers.
package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chat_app/pkg/cookies"
)

type Key string

const (
	UserKey Key = "user"
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
func IsAuth(next http.Handler, secret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := cookies.Get(r, secret)
		if err == nil {
			log.Printf("Cookie found, user %s accepted", user.Username)
			ctxt := context.WithValue(r.Context(), UserKey, user)
			r = r.WithContext(ctxt)
		} else {
			log.Println(err)
		}

		next.ServeHTTP(w, r)
	})
}
