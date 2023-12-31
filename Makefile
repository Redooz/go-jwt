all: build

build:
	@echo "Building..."
	@go build -o main cmd/main.go
# Run the app
run:
	@echo "Running..."
	@go run cmd/main.go
# Test the app
test:
	@echo "Testing..."
	@go test ./...
# Run the app in dev mode
dev:
	@echo "Running in dev mode..."
	@CompileDaemon -command="./main"

.PHONY: all build run test dev