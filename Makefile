build:
	@go build -o bin/gobox

run: build
	@./bin/gobox

test:
	@go test ./... -v

format:
	@go fmt ./...
