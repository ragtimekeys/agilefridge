
all: build

build:
	@go build ./...

clean:
	@rm -rf ./.cache ./server

server: build
	@./server
