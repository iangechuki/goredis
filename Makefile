# Define variables
TARGET := main
BUILD_DIR := build

# Run the built binary
run: build
	./$(BUILD_DIR)/$(TARGET)

# Build the Go project
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(TARGET) main.go
