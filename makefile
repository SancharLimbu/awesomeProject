run:
	swag init -g cmd/main.go
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go