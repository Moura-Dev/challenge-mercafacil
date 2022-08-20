.PHONY: all services clean

down:
	docker-compose down && docker volume prune -f

mod:
	go mod tidy

run:
	go run main.go
services:
	docker-compose up -d