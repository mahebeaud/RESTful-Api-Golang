# Simple Makefile

# Build the application
all: pre-install watch

# Pre-install dependencies and optional tools
pre-install:
	@echo "Installing dependencies"
	@go mod tidy

	@echo "Checking for air installation"
	@if ! command -v ./bin/air > /dev/null; then \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/cosmtrek/air@latest; \
			echo "Air installed."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 0; \
		fi \
	fi

# Build the application
build:
	@echo "Building the application..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@echo "Running the application..."
	@go run cmd/api/main.go

# Test the application - Add test
#test:
#	@echo "Running tests..."
#	@go test ./... -v

# Integration tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/config/database -v

# Clean the binary
clean:
	@echo "Cleaning up..."
	@rm -f main

# Live reload using Air
watch:
	@if command -v ./bin/air > /dev/null; then \
		./bin/air; \
		echo "Starting Air for live reload..."; \
	else \
		echo "Air is not installed. Please install it using 'make pre-install'."; \
	fi

.PHONY: all build run test clean watch itest pre-install
