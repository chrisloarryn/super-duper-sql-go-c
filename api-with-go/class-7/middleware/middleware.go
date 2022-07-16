package middleware

import (
	"crud_psql_7/authorization"
	"github.com/labstack/echo/v4"
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

func Log(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		defer func() {
			elapsed := time.Now().Sub(start).Seconds()
			log.Printf("operation took %g seconds", elapsed)
		}()
		log.Printf("request %q, method: %q", c.Request().URL.Path, c.Request().Method)
		return f(c)
	}
}

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(Authorization)

		if strings.Contains(token, Bearer) {
			str := strings.Split(token, BlankSpace)
			token = strings.TrimSpace(str[1])
		}

		_, err := authorization.ValidateToken(token)
		if err != nil {
			// return response 401
			return c.JSON(http.StatusForbidden, map[string]string{"error": "not allowed"})
		}
		return f(c)
	}
}
