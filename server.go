package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// this is the Router package
	"github.com/gorilla/mux"
)

// Router this initializes the Router
var Router = mux.NewRouter()

func main() {

	// port number
	var port string = ":4000"
	fmt.Println("http://localhost" + port)
	Router.HandleFunc("/", MessagePost).Methods("POST")
	Router.HandleFunc("/create", CreateUser).Methods("POST")
	Router.HandleFunc("/signin", SignIn).Methods("POST")

	Router.HandleFunc("/allusers", AllUsers).Methods("GET")
	Router.HandleFunc("/allcontacts", AllUsers).Methods("GET")
	Router.HandleFunc("/savemsg", SaveMessage).Methods("POST")
	Router.HandleFunc("/addcontact", AddContact).Methods("POST")

	// Websocket connection
	Router.HandleFunc("/socket", SocketHandler)
	go handleMessages()

	// Testing if Server is up
	Router.HandleFunc("/", Test).Methods("GET")

	// Test connect to redis
	// RedisTest()

	// ExampleClient()
	http.ListenAndServe(port, Router)

}

// Test if server is running: GET localhost:4000
func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Go api is running")
}
