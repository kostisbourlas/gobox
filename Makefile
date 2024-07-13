build:
	@go build -o gobox

run: build
	@./gobox

test:
	@go test ./... -v
