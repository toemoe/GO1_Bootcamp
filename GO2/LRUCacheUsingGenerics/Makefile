BINARY_NAME=cache
MAIN_FILE=cache.go
MOD_FILE=go.mod

GO=go
INIT=$(GO) mod init
BUILD=$(GO) build
TEST=$(GO) test
CLEAN=$(GO) clean

all: build test

init:
	@echo "Initializing Go module..."
	$(INIT) cache

test:
	@echo "Running tests..."
	$(TEST) -v

clean:
	@echo "Cleaning..."
	$(CLEAN)
	rm -f $(BINARY_NAME) $(MOD_FILE)
	@echo "Clean complete"

.PHONY: all init build test clean run