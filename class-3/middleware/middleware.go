package middleware

import (
	"log"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
	TokenHash     = "a-entire-securely-token"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			elapsed := time.Now().Sub(start).Seconds()
			log.Printf("operation took %g seconds", elapsed)
		}()
		log.Printf("request %q, method: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Authentication .
func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			elapsed := time.Now().Sub(start).Seconds()
			log.Printf("request %q, method: %q", r.URL.Path, r.Method)
			log.Printf("operation took %g seconds", elapsed)
		}()

		if token := r.Header.Get(Authorization); token != TokenHash {
			// return response 401
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("you do not have authorization to perform the action"))
}
