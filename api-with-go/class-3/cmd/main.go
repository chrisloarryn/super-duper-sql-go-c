package main

import (
	"crud_psql/authorization"
	"crud_psql/handler"
	"crud_psql/storage"
	"fmt"
	"log"
	"net/http"
)

const (
	commonPort = 8080
)

func main() {
	err := authorization.LoadFiles("cmd/certificates/app.rsa", "cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("could not load certificates: %v", err)
	}
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Printf("Server initialized in port: %d", commonPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", commonPort), mux)
	if err != nil {
		log.Printf("error in the server")
	}
}
