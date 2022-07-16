package handler

import (
	"crud_psql_7/middleware"
	"crud_psql_7/storage"
	"github.com/labstack/echo/v4"
)

// RoutePerson .
func RoutePerson(e *echo.Echo, stg *storage.Memory) {
	h := newPerson(stg)

	people := e.Group("/api/v1/people")
	people.Use(middleware.Authentication)

	people.POST("", h.create)
	people.PUT("/:id", h.update)
	people.DELETE("/:id", h.delete)
	people.GET("", h.getAll)
	people.GET("/:id", h.getByID)
}

// RouteLogin .
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/api/v1/login", h.login)
}
