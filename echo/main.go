package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", greeting)
	e.GET("/divide", divide)

	// people group
	//e.POST("/people/create", create)
	//e.GET("/people/read", read)
	//e.PUT("/people/update", update)
	//e.DELETE("/people/delete", delete)

	people := e.Group("/people")
	people.Use(middlewareLogPeople)
	people.POST("", create)
	people.GET("/:id", read)
	people.PUT("/:id", update)
	people.DELETE("/:id", deleter)

	e.Start(":8080")
}

func greeting(c echo.Context) error {
	fmt.Println("sss", c.Request())
	return c.JSON(http.StatusOK, map[string]string{"greeting": "hello, world"})
}

func divide(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)

	if f == 0 {
		return c.String(http.StatusBadRequest, "could not be zero value")
	}

	r := 3000 / f

	return c.String(http.StatusOK, strconv.Itoa(r))
}

func create(c echo.Context) error {
	return c.String(http.StatusCreated, "created")
}

func update(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "updated"+id)
}

func deleter(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "deleted"+id)
}

func read(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "read"+id)
}

func middlewareLogPeople(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("request made to people")
		return f(c)
	}
}
