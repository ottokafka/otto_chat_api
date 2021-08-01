package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: POST
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// We need to use to use the struct model to map the json data to
	type ResponseBody struct {
		Token string `json:"token"`
		User  string `json:"user"`
	}

	var jsonResponse UserSignup

	// We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.Name) // simply print the email

	// Check if user already exists
	user, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	if err != nil {
	}
	println("user ", user)

	check := jsonResponse.Name == user
	if check {

		json.NewEncoder(w).Encode("User " + user + " already exists choose another name")
	} else {

		// Marshal json aka stringify the Json
		userJson, err := json.Marshal(jsonResponse)
		if err != nil {
			println(err)
			return
		}
		errset := RedisClient.HSet("users", jsonResponse.Name, userJson).Err()
		if errset != nil {
			panic(errset)
		}

		json.NewEncoder(w).Encode(ResponseBody{
			Token: CreateToken(jsonResponse.Name),
		})
	}
}
