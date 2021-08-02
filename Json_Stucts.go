package main

type UserResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

// We need to use to use the struct model to map the json data to
type UserSignup struct {
	Name string `json:"name"`
	Pin  int    `json:"pin"`
}
