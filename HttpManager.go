package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func MockApiGo() {
	httpApi, err := ParseService("service.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n",  httpApi)

	maps := map[string]Interaction
	for i, v := range httpApi.Interactions {
		maps.
	}

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
}
