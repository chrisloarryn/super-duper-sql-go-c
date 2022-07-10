package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(msgType, msg string, data interface{}) response {
	return response{
		MessageType: msgType,
		Message:     msg,
		Data:        data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}