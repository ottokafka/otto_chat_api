package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// load from .env file
var err = godotenv.Load()
var secret = os.Getenv("secret")

type UserValid struct {
	User  string `json:"user"`
	Valid bool   `json:"valid"`
}

// Create a token
func CreateToken(username string) string {
	type customClaims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "ottochat.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secureSecretText"))
	if err != nil {
	}
	// fmt.Println(signedToken)

	return signedToken

}

// Authenticate token
func AuthenticateToken(userJwt string) []byte {

	type customClaims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(
		userJwt,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		},
	)
	fmt.Println(err)

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		println("fucked up token")
		var users []UserValid
		users = append(users, UserValid{
			User:  "none",
			Valid: false,
		})

		data, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}
		return data
	}

	// if claims.ExpiresAt < time.Now().UTC().Unix() {
	// 	fmt.Println("jwt Expired", claims.ExpiresAt)
	// }

	username := claims.Username
	// fmt.Println("username ", username)

	// Return json {username: otto, Valid: true}

	var users []UserValid
	users = append(users, UserValid{
		User:  username,
		Valid: true,
	})

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	println("User is Valid returning {user:otto, Valid:true}")
	return jsonUsers
}
