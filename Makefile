build:
	@go build -o bin/fn

run: build
	@./bin/fn

test:
	go test ./... -v