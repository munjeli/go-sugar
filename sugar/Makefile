all: format vet test

vet:
	go vet

format:
	find . -name '*.go' | xargs gofmt -s

build:
	go build

test:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out

