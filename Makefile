build:
	@go mod tidy
	@go build -o bin/habit-tracker

run: build
	@docker compose up -d
	@./bin/habit-tracker

test:
	go test -v
