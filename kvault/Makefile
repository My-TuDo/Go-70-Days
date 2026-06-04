.PHONY:build run-server run-client clean

build:
	go build ./cmd/server/
	go build ./cmd/client/

run-server:
	go run ./cmd/server/

run-client:
	go run ./cmd/client/ $(ARGS)

clean:
	rm -f server client