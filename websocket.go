package main

import (
	// This json package allows us to encode string into a json format

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{} // use default options
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()
	fmt.Println("User Conencted")

	query := r.URL.Query()
	name := query["user"][0]
	fmt.Println(name)

	user, err := Client.HGet("users", name).Result()
	if err != nil {
	}
	fmt.Println("user ", user)

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received from client: %s", message)

		serverMsg := []byte("Hello from api")
		err = conn.WriteMessage(messageType, serverMsg)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
		log.Printf("Send to client: %s", message)
	}
}
