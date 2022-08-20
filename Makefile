.PHONY: all services clean

down:
	docker-compose down && docker volume prune -f

mod:
	go mod tidy

services:
	docker-compose up -d