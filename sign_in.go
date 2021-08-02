package main

import (
	"encoding/json"
	"net/http"
)

// MessagePost: POST
func SignIn(w http.ResponseWriter, r *http.Request) {

	// We need to use to use the struct model to map the json data to

	type ResponseBody struct {
		Token string `json:"token"`
	}

	var jsonResponse UserSignup

	// We decode the incoming data and convert it to a json
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.Name) // simply print the email

	// Check user name and pin number

	// Check if user already exists
	user, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	if err != nil {
		println(err)
	}

	println("Got the user", user)

	var redisJson UserSignup
	bytes := []byte(user)
	err2 := json.Unmarshal(bytes, &redisJson)
	if err2 != nil {
		panic(err2)
	}

	println("user data ", redisJson.Name, redisJson.Pin)

	check := jsonResponse.Name == user
	if check {
		json.NewEncoder(w).Encode("User " + user + " already exists choose another name")
	} else {
		// save user in Redis
		// errset := RedisClient.HSet("users", jsonResponse.Name, jsonResponse.Name).Err()
		// if errset != nil {
		// 	panic(errset)
		// }

		json.NewEncoder(w).Encode(ResponseBody{
			Token: CreateToken(jsonResponse.Name),
		})
	}
}
