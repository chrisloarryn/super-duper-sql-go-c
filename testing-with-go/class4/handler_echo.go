package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// PersonWithEcho estructura de una persona
type PersonWithEcho struct {
	Name string
	Age  int
}

// GetWithEcho handler que retorna una persona
func GetWithEcho(c echo.Context) error {
	p := Person{
		"Jhoana",
		31,
	}

	return c.JSON(http.StatusOK, p)
}
