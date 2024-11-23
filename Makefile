run:
	@echo "Running server..."
	go run cmd/server/main.go

tidy:
	@echo "Tidying up..."
	go mod tidy
	go fmt ./..

.PHONY: run tidy

