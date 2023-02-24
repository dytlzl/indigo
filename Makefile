lint:
	staticcheck ./...

fmt:
	goimports -w .

mod:
	go mod tidy

test:
	go test ./...

gen:
	go generate ./...

all: gen fmt lint test