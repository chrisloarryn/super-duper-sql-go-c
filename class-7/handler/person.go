package handler

import (
	"crud_psql_7/model"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(s Storage) person {
	return person{s}
}

func (p *person) create(c echo.Context) error {
	var data model.Person
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person has not right structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "something went wrong", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {
	if id := c.Param("id"); len(id) == 0 {
		response := newResponse(Error, "id is necessary for updating a person", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "id should be a number positive", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var data model.Person
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person has not right structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(id, &data)
	if err != nil {
		fmt.Printf("data :: %+v", data)
		response := newResponse(Error, "something went wrong", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "person details has been updated", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {
	if id := c.Param("id"); len(id) == 0 {
		response := newResponse(Error, "id is necessary for deleting a person", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "id should be a number positive", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(id)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "that id does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)

	}

	if err != nil {
		response := newResponse(Error, "an error has occurred when deleting resource", nil)
		return c.JSON(http.StatusInternalServerError, response)

	}

	response := newResponse(Message, "Ok", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {
	resp, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "something went wrong", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	if len(resp) == 0 {
		response := newResponse(Error, "", nil)
		return c.JSON(http.StatusNoContent, response)
	}

	response := newResponse(Message, "Ok", resp)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {
	if id := c.Param("id"); len(id) == 0 {
		response := newResponse(Error, "id is necessary for deleting a person", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "id should be a number positive", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := p.storage.GetByID(id)
	if err != nil {
		response := newResponse(Error, "resource was not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "that id does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err != nil {
		response := newResponse(Error, "an error has occurred when fetching resource", nil)
		return c.JSON(http.StatusInternalServerError, response)

	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
