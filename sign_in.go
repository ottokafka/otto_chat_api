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
	println("incoming user data ", jsonResponse.Name, jsonResponse.Pin)
	// Check if user and password before giving a token
	redisUser, err := RedisClient.HGet("users", jsonResponse.Name).Result()
	if err != nil {
		println(err)
	}

	// fmt.Println("Got the user", redisUser)

	// Unmarshal the Json from Redis
	var redisJson UserSignup
	bytes := []byte(redisUser)
	err2 := json.Unmarshal(bytes, &redisJson)
	if err2 != nil {
		panic(err2)
	}

	println("Redis user data ", redisJson.Name, redisJson.Pin)
	// checkName := jsonResponse.Name != redisJson.Name

	if redisJson.Pin == jsonResponse.Pin && jsonResponse.Name == redisJson.Name {
		println("user and password match send a tooken")
		// Password is correct send him/her a token brother
		json.NewEncoder(w).Encode(ResponseBody{
			Token: CreateToken(jsonResponse.Name),
		})
	} else {
		println("Not matching user and password ")
		// user pass is wrong send back message
		json.NewEncoder(w).Encode(UserResponse{
			Message: "Invalid Credentials",
		})
	}
}
