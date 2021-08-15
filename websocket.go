package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	Message string `json:"message"`
	Date    string `json:"date"`
	User    string `json:"user"`
	User2   string `json:"user2"`
}

type Users struct {
	User  string `json:"user"`
	User2 string `json:"user2"`
}

// var names = make(map[int]string)
// var users = make(map[string]string)

// var clients = make(map[*websocket.Conn]bool)
var clients = make(map[string]*websocket.Conn)
var broadcaster = make(chan ChatMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	// AUthorization for later use

	// token := r.Header.Values("Authorization")
	// var user = AuthenticateToken(token[0])
	user := r.Header.Values("user")[0]
	user2 := r.Header.Values("user2")[0]

	println("user & user2 ", user, user2)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()

	// users[user] = user
	// fmt.Println("users list", users)

	// println(users)

	// Add the websocket address to clients list
	// clients[ws] = true
	clients[user] = ws

	fmt.Println("clients print all", clients)

	// if it's zero, no messages were ever sent/saved
	var flipped = false
	if RedisClient.Exists(user+":"+user2).Val() != 0 {
		sendPreviousMessages(ws, user, user2, flipped)
	} else if RedisClient.Exists(user2+":"+user).Val() != 0 {
		flipped = true
		sendPreviousMessages(ws, user2, user, flipped)
	} else {
		println("user & user2 ", user, user2, "are new users")
		// start handling messages
		// go handleMessages(user, user2)
	}

	// Ever message will be passed through here
	for {
		var msg ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, user)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}

func sendPreviousMessages(ws *websocket.Conn, user string, user2 string, flipped bool) {

	println("sendPreviousMessages", user)
	// Grabs messages from Redis List
	chatMessages, err := RedisClient.LRange(user+":"+user2, 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// send previous messages
	for _, chatMessage := range chatMessages {
		var msg ChatMessage
		json.Unmarshal([]byte(chatMessage), &msg)
		// print(chatMessage)
		if flipped {
			messageClient(user2, msg)
		} else {
			messageClient(user, msg)
		}
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {

	for {
		// grab any next message from channel
		msg := <-broadcaster
		println()
		// usersCurrent := <-userchannel
		// log.Println("users CUrrent", usersCurrent)

		// check storage before saving
		if RedisClient.Exists(msg.User+":"+msg.User2).Val() != 0 {
			println("saving ", msg.User+":"+msg.User2)
			storeInRedis(msg, msg.User, msg.User2)
		} else if RedisClient.Exists(msg.User2+":"+msg.User).Val() != 0 {
			println("saving ", msg.User2+":"+msg.User)
			storeInRedis(msg, msg.User2, msg.User)
		} else {
			storeInRedis(msg, msg.User, msg.User2)
		}

		log.Println("message from user:", msg.User, msg.User2)
		//
		messageClients(msg)

	}
}

func storeInRedis(msg ChatMessage, user string, user2 string) {
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	// Save messages to Redis list
	println("saving: ", user+":"+user2)
	if err := RedisClient.RPush(user+":"+user2, json).Err(); err != nil {
		panic(err)
	}
}

func messageClients(msg ChatMessage) {
	// send to every client currently connected
	for user := range clients {
		// Get the websocketAddresses
		websocketAddress, userOnline := clients[user]
		websocketAddress2, user2Online := clients[msg.User2]
		println("current client", user, "websocketAddress", websocketAddress, "isMapContainsKey", userOnline)
		fmt.Println("All clients map", clients)

		if user == msg.User {
			fmt.Println("Send back to user1")
			messageClient(user, msg)
			if user2Online {
				println("User2 is online", user2Online, websocketAddress2)
				messageClient(msg.User2, msg)
				break
			} else {
				println("User2 is not online", websocketAddress2)
			}
			break
		}
	}
}

func messageClient(user string, msg ChatMessage) {
	fmt.Println(clients, user)
	websocketAddress := clients[user]
	err := websocketAddress.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		websocketAddress.Close()
		// delete(clients, client)
	}

}
