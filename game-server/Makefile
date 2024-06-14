# Variables
GO=go
PROJECT_NAME=game-server
BUILD_DIR=bin

# Build the project
.PHONY: all build
all: build

build:
	$(GO) build -o $(BUILD_DIR)/$(PROJECT_NAME) .

# Run the project
.PHONY: run
run:
	$(BUILD_DIR)/$(PROJECT_NAME)

# Clean the build directory
.PHONY: clean
clean:
	if exist $(BUILD_DIR) (rmdir /s /q $(BUILD_DIR))
