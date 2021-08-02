package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// MessagePost: POST
func AddContact(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Values("Authorization")
	fmt.Println(len(token))
	if len(token) == 0 {
		json.NewEncoder(w).Encode(UserResponse{
			Message: "You dont have a token, please sign in",
			User:    "",
		})

		// Check if token is Valid
	} else {
		var tokenUser = AuthenticateToken(token[0])

		var users []UserValid
		err := json.Unmarshal([]byte(tokenUser), &users)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("This token belong to: " + users[0].User)
		fmt.Println("Is a valid token?: ", users[0].Valid)

		// Token is Valid continue
		if users[0].Valid == true {

			// Get select username the user wants to add from incoming json body
			var jsonResponse UserResponse

			// We decode the incoming data and convert it to a json this gets sent to the client
			json.NewDecoder(r.Body).Decode(&jsonResponse)
			fmt.Println(jsonResponse)
			println("This user is search this name: ", jsonResponse.User)

			// Check if the user exists
			username, err := RedisClient.HGet("users", jsonResponse.User).Result()
			// Error means user doesnt exist return response
			if err != nil {
				println("user not here", jsonResponse.User)
				json.NewEncoder(w).Encode(UserResponse{
					Message: "This " + jsonResponse.User + " is not a user, Maybe you typed the name wrong",
					User:    "",
				})
			} else {
				println("User is in the database ", username)

				// Send back user name
				json.NewEncoder(w).Encode(UserResponse{
					Message: "user is in the system",
					User:    jsonResponse.User,
				})
			}

		} else {
			// return not valid token
			json.NewEncoder(w).Encode(UserResponse{
				Message: "Token is invalid, please sign in again",
				User:    "",
			})
		}
	}

}
