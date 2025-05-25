tidy:
	@go mod tidy

build:	
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

clean:
	@rm -rf bin
	@echo "Cleaned up bin directory"