.PHONY: dev build lint

dev:
	@echo "Starting development server..."
	@GO_ENV=development go run main.go
build:
	@echo "Building binary..."
	@go build -o app main.go
lint:
	@golangci-lint run
