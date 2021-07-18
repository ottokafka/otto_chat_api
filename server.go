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
	Router.HandleFunc("/allusers", AllUsers).Methods("GET")
	Router.HandleFunc("/savemsg", SaveMessage).Methods("POST")

	// Websocket connection
	Router.HandleFunc("/socket", SocketHandler)

	// Testing if Server is up
	Router.HandleFunc("/", Test).Methods("GET")

	// Test connect to redis
	RedisTest()

	http.ListenAndServe(port, Router)
}

func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Go api is running")
}
