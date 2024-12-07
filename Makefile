# Define variables
APP_NAME := goredis
GO_FILES = $(shell find . -type f -name "*.go" -not -name '*_test.go')
BUILD_DIR := build

# Default target
all: run

#Run the application
run:
	go run $(GO_FILES)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

# Clean the build artifacts
clean:
	rm -rf $(BUILD_DIR)

#format the code
fmt:
	go fmt ./...

.PHONY: all run build test clean fmt
