package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type HttpEndpoint struct {
	Interactions []Interaction
}

type Interaction struct {
	Request Request
}

type Request struct {
	Method string
	Path string
}

type ServiceParser struct {}

func ParseService(filename string) (HttpEndpoint, error) {
	body, readFileErr := ioutil.ReadFile("service.json")
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var endpoint HttpEndpoint
	err := json.Unmarshal(body, &endpoint)

	return endpoint, err
}