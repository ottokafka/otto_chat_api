package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: GET
func LoadMessages(w http.ResponseWriter, r *http.Request) {

	// We need to use to use the struct model to map the json data to
	type User struct {
		Name string `json:"name"`
		Pin  int    `json:"pin"`
	}

	type responseBody struct {
		Messages string `json:"token"`
		User     string `json:"user"`
	}

	var jsonResponse User

	// We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.Name) // simply print the email

	// Check if user exists
	user, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	if err != nil {
	}

	check := jsonResponse.Name == user
	if !check {
		json.NewEncoder(w).Encode("User " + user + " has no messages")
	} else {
		// load user msg from Redis
		msgs, err := RedisClient.HGet("users", jsonResponse.Name).Result()
		if err != nil {
		}
		println("load msgs", msgs)

		json.NewEncoder(w).Encode(responseBody{
			Messages: "example-auth-token",
			User:     jsonResponse.Name,
		})
	}

}
