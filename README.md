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
- CREATE MOCKS: `~/.asdf/installs/golang/1.15.2/packages/bin/mockgen -destination=mocks/mock_database_connection.go -package=mocks -source=core_module/frameworks/database/connection.go` `https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f`