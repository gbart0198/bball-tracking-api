build:
	@go build -o bin/bball-tracker-api

run: build
	@./bin/bball-tracker-api

test:
	@go test -v ./...
