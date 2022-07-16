package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	url = "http://localhost:8080"
)

func main() {
	lc := loginClient(url+"/api/v1/login", "contact@chrisloarryn.cl", "123456789")

	p := Person{
		Name: "Alexys Lozada",
		Age:  42,
		Communities: Communities{
			{
				Name: "EDteam",
			},
			{
				Name: "Golang in Spanish",
			},
		},
	}
	gr := createPerson(url+"/api/v1/people", lc.Data.Token, &p)

	fmt.Printf("THE RESPONSE:: %+v", gr)
}

func httpClient(method, u, t string, b io.Reader) *http.Response {
	req, err := http.NewRequest(method, u, b)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", t)

	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	return response
}
