package main

// import (
// 	// This json package allows us to encode string into a json format
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// type ResponseBody struct {
// Message string `json:"message"`
// Date    string `json:"date"`
// User    string `json:"user"`
// User2   string `json:"user2"`
// }

// var clients = make(map[*websocket.Conn]bool)
// var broadcaster = make(chan ResponseBody)
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func SocketHandler(w http.ResponseWriter, r *http.Request) {
// 	token := r.Header.Values("Authorization")
// 	var tokenUser = AuthenticateToken(token[0])

// 	user, err := RedisClient.HGet("users", tokenUser).Result()
// 	if err != nil {
// 	}
// 	authUser := user == tokenUser
// 	// print(authUser)

// 	if authUser {
// 		print("user authenticated")

// 		upgrader = websocket.Upgrader{} // use default options
// 		// Upgrade our raw HTTP connection to a websocket based one
// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Print("Error during connection upgradation:", err)
// 			return
// 		}
// 		defer conn.Close()
// 		clients[conn] = true

// 		fmt.Println("User Conencted")

// 		// The event loop
// 		for {

// 			// Read in a new message as JSON and map it to a Message object

// 			messageType, message, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Println("Error during message reading:", err)
// 				break
// 			}
// 			// log.Printf("Received from client: %s", message)

// 			// decode json message
// 			var jsonResponse ResponseBody

// 			json.Unmarshal([]byte(message), &jsonResponse)

// 			fmt.Printf("%+v\n", jsonResponse)

// 			// check if user2 is in redis
// 			user2, err := RedisClient.HGet("users", jsonResponse.User2).Result()
// 			if err != nil {
// 				println(user2 + "user2 is not found in redis")
// 			}

// 			var checkUser2 bool = user2 == jsonResponse.User2

// 			if checkUser2 {

// 				errset := RedisClient.HSet("messages", jsonResponse.User, message).Err()
// 				if errset != nil {
// 					panic(errset)
// 				}

// 				serverMsg := []byte("Server: " + string(message))

// 				err = conn.WriteMessage(messageType, serverMsg)
// 				if err != nil {
// 					log.Println("Error during message writing:", err)
// 					// break
// 				}
// 				// log.Printf("Send to client: %s", message)
// 			} else {
// 				println("User2 is not a user")
// 			}
// 		}
// 	} else {
// 		print("User is not sign up or token is not correct")
// 		json.NewEncoder(w).Encode("user not in the system")
// 	}
// }
