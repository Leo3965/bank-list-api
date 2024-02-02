build:
	@go build -o bin/bank-list-api

run: build
	@./bin/bank-list-api

test:
	@go test -v ./...

