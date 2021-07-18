package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: POST
func MessagePost(w http.ResponseWriter, r *http.Request) {

	// We need to use to use the struct model to map the json data to
	type Business struct {
		Msg   string `json:"msg"`
		Phone int    `json:"phone"`
	}

	var jsonResponse Business

	// We decode the incoming data and convert it to a json this gets sent to the client
	json.NewDecoder(r.Body).Decode(&jsonResponse)
	println("Message to Server", jsonResponse.Msg) // simply print the email
	json.NewEncoder(w).Encode("post sent Successfully")

}
