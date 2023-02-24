lint:
	staticcheck ./...

fmt:
	goimports -w .