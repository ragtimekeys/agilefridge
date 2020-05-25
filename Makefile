
all: build

build: generate-openapi-server
	@go build ./...

generate-openapi-server:
	@./build/openapi-generator generate \
	  -i ./openapi.yml \
	  -g go-server \
	  -o ./internal/openapi

clean:
	@rm -rf ./.cache ./server ./internal/openapi

server: build
	@./server
