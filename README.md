# GO LANG - REST API

This simple project has the purpose to study Go Lang, Clean Architecture and REST API.

## Features:

- Create User Account
- Create User Sessions with JWT
- Save the record of every time the user login
- Manager User Control Access (ACL) with UsersPermissions


### utils

- RUN: `go run src/*.go`
- BUILD: `go build src/*.go`
- TEST ALL FILES: `go test ./src/...`
- TEST COV: `go test ./src/... -cover`
- TEST COV: `go test ./src/... -cover -coverprofile=c.out && go tool cover -html=c.out -o coverage.html`