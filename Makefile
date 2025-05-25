tidy:
	@go mod tidy
	@clear

build:	
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: tidy build
	@./bin/ecom

# clean bin directory
clean:
	@rm -rf bin
	@echo "Cleaned up bin directory"

# create sql migration files
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# create database tables
migrate-up:
	@go run cmd/migrate/main.go up

# remove database tables
migrate-down:
	@go run cmd/migrate/main.go down