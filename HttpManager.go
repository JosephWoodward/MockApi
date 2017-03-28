package main

import (
	"fmt"
	"net/http"
)

var httpMaps = make(map[string]Interaction)

func handler(w http.ResponseWriter, r *http.Request) {


	url := r.URL.Path[1:]
	fmt.Fprintf(w, "Hi there, I love %s!", url)

}

func MockApiGo() {
	httpApi, err := ParseService("service.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n",  httpApi)

	for _, interaction := range httpApi.Interactions {
		fmt.Println(interaction.Request.Path)

		httpMaps[interaction.Request.Path] = interaction
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
