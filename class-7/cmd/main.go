package main

import (
	"crud_psql_7/authorization"
	"crud_psql_7/handler"
	"crud_psql_7/middleware"
	"crud_psql_7/storage"
	"github.com/labstack/echo/v4"
	"log"
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

	e := echo.New()
	e.Use(middleware.Log)

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Printf("Server initialized in port: %d", commonPort)

	if err = e.Start(":8080"); err != nil {
		log.Printf("error in the server")
	}
}
