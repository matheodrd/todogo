build:
	@go build -o bin/todogo

run: build
	@./bin/todogo
