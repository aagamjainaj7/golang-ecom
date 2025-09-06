# Build the Go project
build:
	@go build -o bin/ecom "./cmd/main.go"

# Run tests
test:
	@go test -v ./...

# Run the app (depends on build)
run: build
	@./bin/ecom
