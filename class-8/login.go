package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error in marshal of login %v", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("close")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("want: %d, but got: %d, resp: %s", http.StatusOK, resp.StatusCode, body)
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error unmarshal in body login: %v", err)
	}

	return dataResponse
}
