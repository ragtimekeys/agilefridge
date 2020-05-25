
all: build

build: generate-openapi-server
	@go build ./...

generate-openapi-server:
	@./build/openapi-generator generate \
	  -i ./openapi.yml \
	  -g go-server \
	  -o ./.cache/openapi-server-generated

clean:
	@rm -rf ./.cache ./server

server: build
	@./server
