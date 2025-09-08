# Build the Go project
build:
	@go build -o bin/ecom "./cmd/main.go"

# Run tests
test:
	@go test -v ./...

# Run the app (depends on build)
run: build
	@./bin/ecom

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down