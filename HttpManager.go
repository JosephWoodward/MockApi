package main

import (
	"fmt"
	"net/http"
	"strings"
)

var httpMaps = make(map[string]Interaction)
var writeError = func(w http.ResponseWriter, request *http.Request) {
	http.Error(w, "Endpoint not found", http.StatusNotFound)
}

func handler(w http.ResponseWriter, request *http.Request) {

	// TODO: Tidy up
	// Path exists
	path := request.URL.Path[0:]
	interaction, exists := httpMaps[path]
	if !exists {
		writeError(w, request)
		return
	}

	// Correct method
	if strings.ToLower(request.Method) != strings.ToLower(interaction.Request.Method) {
		writeError(w, request)
		return
	}

	// Correct header accept type
	//header := request.Header.Get("accept")
	//if strings.ToLower(header) == "application/json" {
	//	writeError(w, request)
	//	return
	//}

	for _, header := range interaction.Response.Headers  {
		_ = header
	}

	// Write response
	w.Header().Add("status", "200")
	w.Write([]byte("Testing"))
}

func MockApiGo(filename string) {
	httpApi, err := ParseService(filename)
	if err != nil {
		fmt.Println(err)
	}

	for _, interaction := range httpApi.Interactions {
		_, exists := httpMaps[interaction.Request.Path]
		if !exists {
			httpMaps[interaction.Request.Path] = interaction
		} else {
			panic("Duplicate keys")
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}