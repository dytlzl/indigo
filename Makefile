gen:
	go generate ./...

fmt:
	goimports -w .

mod:
	go mod tidy

test:
	go test ./...

lint:
	golangci-lint run

all: gen fmt mod test lint