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

tailwindcss-generate:
	@echo "Generating temple..."
	@npx @tailwindcss/cli -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css