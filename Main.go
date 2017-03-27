package main

import (
	"fmt"
	"encoding/json"
)

func main(){

	api, parseErr := ParseService("service.json")

	if parseErr != nil {
		fmt.Println("Could not read file\n", parseErr)
	}

	var endpoint HttpEndpoint
	err := json.Unmarshal([]byte(api), &endpoint)
	if err == nil {
		fmt.Printf("%+v\n",  endpoint)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n",  endpoint)
	}

}