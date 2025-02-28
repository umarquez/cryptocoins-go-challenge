BUILD_DIR := ./bin
APP_NAME := app

.PHONY: build clean test

setup:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest

all: clean setup generate test build

generate: setup
	go generate ./...
	swag init -g main.go -d cmd/app -o docs/

build: generate
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/app/main.go

clean:
	rm -rf $(BUILD_DIR)
	rm -rf ./doc
	rm -rf ./data

test:
	go test -v ./...

run: setup clean build
	chmod +x $(BUILD_DIR)/$(APP_NAME)
	ls -l $(BUILD_DIR)/$(APP_NAME)
	$(BUILD_DIR)/$(APP_NAME)