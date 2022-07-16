package handler

import (
	"crud_psql/authorization"
	"crud_psql/model"
	"encoding/json"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "method now allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "not valid struct", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "user or password not valid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "could not generate token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}

	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contact@chrisloarryn.cl" && data.Password == "123456789"
}
