package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: GET
func SaveMessage(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Values("Authorization")

	// // We need to use to use the struct model to map the json data to
	// type User struct {
	// 	Name string `json:"name"`
	// }

	type ResponseBody struct {
		Messages string `json:"message"`
		Time     string `json:"time"`
		User     string `json:"user"`
		User2    string `json:"user2"`
	}

	var jsonResponse ResponseBody

	// // We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	AuthenticateToken(token[0])

	println("save message user: ", jsonResponse.User) // simply print the email

	// // Check if user exists
	// user, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	// if err != nil {
	// }

	// check := jsonResponse.Name == user
	// if !check {
	// 	json.NewEncoder(w).Encode("User " + user + " has no messages")
	// } else {
	// 	// load user msg from Redis
	// 	msgs, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	// 	if err != nil {
	// 	}
	// 	println("load msgs", msgs)

	// 	json.NewEncoder(w).Encode(responseBody{
	// 		Messages: "example-auth-token",
	// 		User:     jsonResponse.Name,
	// 	})
	// }

}
