package main

import (
	"net/http"
)

// MessagePost: GET
func SaveMessage(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Values("Authorization")

	// // We need to use to use the struct model to map the json data to
	// type User struct {
	// 	Name string `json:"name"`
	// }

	// type responseBody struct {
	// 	Messages string `json:"token"`
	// 	User     string `json:"user"`
	// }

	// type msgBody struct {
	// 	Msg  string `json:"msg"`
	// 	Time string `json:"time"`
	// }

	// var jsonResponse User

	// // We decode the incoming data and convert it to a json
	// json.NewDecoder(r.Body).Decode(&jsonResponse)

	AuthenticateToken(token[0])

	// println("Message to Server", jsonResponse.Name) // simply print the email

	// // Check if user exists
	// user, err := Client.HGet("users", jsonResponse.Name).Result()
	// if err != nil {
	// }

	// check := jsonResponse.Name == user
	// if !check {
	// 	json.NewEncoder(w).Encode("User " + user + " has no messages")
	// } else {
	// 	// load user msg from Redis
	// 	msgs, err := Client.HGet("users", jsonResponse.Name).Result()
	// 	if err != nil {
	// 	}
	// 	println("load msgs", msgs)

	// 	json.NewEncoder(w).Encode(responseBody{
	// 		Messages: "example-auth-token",
	// 		User:     jsonResponse.Name,
	// 	})
	// }

}
