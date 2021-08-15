package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MessagePost: GET
func AllContacts(w http.ResponseWriter, r *http.Request) {
	// TODO: Check token here

	// query := r.URL.Query()
	// name := query["user"][0]
	// fmt.Println(name)

	var jsonResponse UserResponse

	// Check the user

	// We decode the incoming data and convert it to a json this gets sent to the client
	json.NewDecoder(r.Body).Decode(&jsonResponse)

	println("Message to Server", jsonResponse.User) // simply print the email

	all, err := RedisClient.HGetAll("names").Result()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(all)

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = RedisClient.Scan(cursor, "names", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}

}
