package middleware

import (
	"crud_psql/authorization"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	Authorization = "Authorization"
	Bearer        = "Bearer"
	BlankSpace
	TokenHash = "a-entire-securely-token"
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

		token := r.Header.Get(Authorization)

		if strings.Contains(token, Bearer) {
			str := strings.Split(token, BlankSpace)
			token = strings.TrimSpace(str[1])
		}

		_, err := authorization.ValidateToken(token)
		fmt.Println(Authorization, err)
		if err != nil {
			// return response 401
			forbidden(w, r, err.Error())
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(fmt.Sprintf(`{"message_type": "error" , "message": "%s", "data": null}`, msg)))
}
