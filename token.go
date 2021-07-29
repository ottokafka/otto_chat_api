package main

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// load from .env file
var err = godotenv.Load()
var secret = os.Getenv("secret")

// Create a token
func CreateToken(username string) string {
	type customClaims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "macchat.com",
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
func AuthenticateToken(userJwt string) string {

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

	}

	// if claims.ExpiresAt < time.Now().UTC().Unix() {
	// 	fmt.Println("jwt Expired", claims.ExpiresAt)
	// }

	username := claims.Username
	// fmt.Println("username ", username)

	return username
}
