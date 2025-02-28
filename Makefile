BUILD_DIR := bin
APP_NAME := app

.PHONY: build clean test

setup:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest

all: build

generate:
	go generate ./...
	swag init -d cmd/app/ -o doc/

build: generate test
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/app/main.go

clean:
	rm -rf $(BUILD_DIR)
	rm -f ./doc

test:
	go test -v ./...