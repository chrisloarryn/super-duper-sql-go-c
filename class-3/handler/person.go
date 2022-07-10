package handler

import (
	"crud_psql/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(s Storage) person {
	return person{s}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "method now allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	var data *model.Person
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "person has not right structure", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	fmt.Printf("%+v", data)

	err = p.storage.Create(data)
	if err != nil {
		response := newResponse(Error, "something went wrong", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "method now allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "id should be a number positive", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	var data *model.Person
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "person has not right structure", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(id, data)
	if err != nil {
		response := newResponse(Error, "something went wrong", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "person details has been updated", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "method now allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "id should be a number positive", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(id)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "that id does not exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := newResponse(Error, "an error has occurred when deleting resource", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "method now allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	resp, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "something went wrong", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	if len(resp) == 0 {
		response := newResponse(Error, "", nil)
		responseJSON(w, http.StatusNoContent, response)
		return
	}

	response := newResponse(Message, "Ok", resp)
	responseJSON(w, http.StatusOK, response)
}
