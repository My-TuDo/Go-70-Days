.PHONY: build run clean

build:
	go build ./cmd/api/

run:
	go run ./cmd/api/

clean:
	rm -f api

up:
	docker-compose up -d --build

down:
	docker-compose down