package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: GET
func LoadMessages(w http.ResponseWriter, r *http.Request) {

	var jsonResponse UserResponse

	// We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.User) // simply print the email

	// Check if user exists
	user, err := RedisClient.HGet("users", jsonResponse.User).Result()
	if err != nil {
	}

	check := jsonResponse.User == user
	if !check {
		json.NewEncoder(w).Encode("User " + user + " has no messages")
	} else {
		// load user msg from Redis
		msgs, err := RedisClient.HGet("users", jsonResponse.User).Result()
		if err != nil {
		}
		println("load msgs", msgs)

		json.NewEncoder(w).Encode(UserResponse{
			Message: "Load your messages",
			User:    jsonResponse.User,
		})
	}

}
