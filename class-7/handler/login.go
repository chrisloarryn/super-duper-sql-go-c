package handler

import (
	"crud_psql_7/authorization"
	"crud_psql_7/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := &model.Login{}
	err := c.Bind(data)
	if err != nil {
		fmt.Printf("my error %s", err.Error())
		resp := newResponse(Error, "not valid struct", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(data) {
		resp := newResponse(Error, "user or password not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(data)
	if err != nil {
		resp := newResponse(Error, "could not generate token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}

	resp := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contact@chrisloarryn.cl" && data.Password == "123456789"
}
