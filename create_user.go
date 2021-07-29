package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: POST
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// We need to use to use the struct model to map the json data to
	type User struct {
		Name string `json:"name"`
		Pin  int    `json:"pin"`
	}

	type responseBody struct {
		Token string `json:"token"`
	}

	var jsonResponse User

	// We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.Name) // simply print the email

	// Check if user already exists
	user, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	if err != nil {
	}

	check := jsonResponse.Name == user
	if check {
		json.NewEncoder(w).Encode("User " + user + " already exists choose another name")
	} else {
		// save user in Redis
		errset := RedisClient.HSet("users", jsonResponse.Name, jsonResponse.Name).Err()
		if errset != nil {
			panic(errset)
		}

		json.NewEncoder(w).Encode(responseBody{
			Token: CreateToken(jsonResponse.Name),
		})
	}
}
