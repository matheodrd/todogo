build:
	@go build -o bin/gotodo

run: build
	@./bin/gotodo