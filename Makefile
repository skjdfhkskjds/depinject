tidy:
	go mod tidy

test-unit:
	@echo "Running unit tests..."
	@go test -v $(shell go list ./... | grep -v '/examples')

# Run all main.go files in examples/ and ensure they don't exit unexpectedly.
test-integration:
	@echo "Running integration tests..."
	@go test ./examples/...
