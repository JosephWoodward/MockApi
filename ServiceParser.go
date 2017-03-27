package main

import (
	"io/ioutil"
	"encoding/json"
)

type service []byte


type HttpEndpoint struct {
	Uri string
	Name string
}

type ServiceParser struct {}

func ParseService(filename string) ([]HttpEndpoint, error) {
	body, readFileErr := ioutil.ReadFile("service.json")

	var endpoint HttpEndpoint
	err := json.Unmarshal([]byte(body), &endpoint)

	return body, readFileErr
}