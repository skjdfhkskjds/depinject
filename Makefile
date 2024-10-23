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

# Run a specific example based on the directory name if it exists.
run:
	@if [ -n "$(example)" ]; then \
		main="$(EXAMPLES_DIR)/$(example)/main.go"; \
		if [ -f "$$main" ]; then \
			echo "Running $$main..."; \
			go run $$main || { echo "$$main failed"; exit 1; }; \
		fi; \
	else \
		echo "Please specify an example directory using 'example' variable"; \
	fi
