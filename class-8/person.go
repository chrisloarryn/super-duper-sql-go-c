package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func createPerson(url, t string, p *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(p)
	if err != nil {
		log.Fatalf("person marshal error: %v", err)
	}

	resp := httpClient(http.MethodPost, url, t, data)
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

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("want: %d, but got: %d, resp: %s", http.StatusCreated, resp.StatusCode, body)
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error unmarshal in body login: %v", err)
	}

	return dataResponse
}
