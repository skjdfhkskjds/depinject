tidy:
	go mod tidy

test-unit:
	@echo "Running unit tests..."
	@go list -f '{{.Dir}}/...' -m | xargs \
		go test

EXAMPLES_DIR := ./examples
MAIN_FILES := $(shell find $(EXAMPLES_DIR) -name main.go)

# Run all main.go files in examples/ and ensure they don't exit unexpectedly.
test-examples:
	@for main in $(MAIN_FILES); do \
		echo "Running $$main..."; \
		go run $$main || { echo "$$main failed"; exit 1; }; \
	done