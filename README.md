# GETTING STARTED creating a fresh GO project

```bash
mkdir hello
cd hello
go mod init example.com/hello
```

# Download live reload with air

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
