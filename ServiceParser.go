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
	Response Response
}

type Request struct {
	Method string
	Path string
	Headers map[string]Header
}

type Header map[string]HeaderItem

type HeaderItem struct {
	ContentType string `json:"contentType"`
}

type Response struct {
	Status int
	Headers Header
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