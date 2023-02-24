lint:
	staticcheck ./...

fmt:
	goimports -w .

mod:
	go mod tidy

test:
	go test ./...

all: fmt lint test