package handler

import (
	"crud_psql/middleware"
	"crud_psql/storage"
	"net/http"
)

// RoutePerson .
func RoutePerson(mux *http.ServeMux, stg *storage.Memory) {
	h := newPerson(stg)

	mux.HandleFunc("/api/v1/people/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/api/v1/people/update", middleware.Log(h.update))
	mux.HandleFunc("/api/v1/people/delete", middleware.Log(h.delete))
	mux.HandleFunc("/api/v1/people/get-all", middleware.Log(h.getAll))
}
