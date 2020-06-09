all: test

build:
	go build

test:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out

