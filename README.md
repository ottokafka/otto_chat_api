# Chat api

This is a simple chat api built with the go language. 
Connects to redis to save messages
use JWT method to authenticate users

# Packages used

go get github.com/joho/godotenv

This is basically like node to grab .env file


# GETTING STARTED

Need to have go installed ( macOs: `brew install go`)

`go run .` 

Want auto reload and auto rebuild read below - air


# creating a fresh GO project

```bash
mkdir hello
cd hello
go mod init example.com/hello
```

# Download live reload with air ( kinda like nodemon )

```
# binary will be $(go env GOPATH)/bin/air

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

code ~/.zshrc
# add this line 
export PATH=$PATH:$(go env GOPATH)/bin
```

Inside the project directory you want to live reload

```
air init
air
```
