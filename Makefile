# Simple Makefile for a Go project

build:
	@echo "Building..."

	 @go build -o build/main cmd/api/main.go
# Run the application
run:
	@go run cmd/api/main.go

temple-generate:
	@echo "Generating temple..."
	@templ generate