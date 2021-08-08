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
	if RedisClient.HExists("users", jsonResponse.Name).Val() == true {
		println("a user does indeed exist")
		println(RedisClient.HExists("users", jsonResponse.Name).Val())
		// send Json telling him name already exists bro
		json.NewEncoder(w).Encode("User " + jsonResponse.Name + " already exists choose another name")

	} else {
		println("no user exist go ahead and create that mother fucka")

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

// // check if name in redis if so tell him no go
// if jsonResponse.Name == redisJson.Name {

// 	json.NewEncoder(w).Encode("User " + redisJson.Name + " already exists choose another name")
// } else {

// }

// redisUser, err := RedisClient.HGet("users", jsonResponse.Name).Result()
// if err != nil {
// 	println(err)
// }

// fmt.Println(redisUser)
// // Unmarshal the Json from Redis
// var redisJson UserSignup
// bytes := []byte(redisUser)
// err2 := json.Unmarshal(bytes, &redisJson)
// if err2 != nil {
// 	panic(err2)
// }
