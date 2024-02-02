build:
	@go build -o bin/bank-list-api

run: build
	@./bin/bank-list-api

test:
	@go test -v ./...

create-postgres-image:
	 docker run --name banklistdb -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -e POSTGRES_DB=banklistdb -p 5432:5432 -d postgres
