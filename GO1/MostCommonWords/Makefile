BINARY_NAME=wordcounter
MAIN_FILE=main.go
MOD_FILE=go.mod

GO=go
INIT=$(GO) mod init
BUILD=$(GO) build
TEST=$(GO) test
CLEAN=$(GO) clean

all: build test

init:
	@echo "Initializing Go module..."
	$(INIT) mostcommonwords

build:
	@echo "Building $(BINARY_NAME)..."
	$(BUILD) -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete: ./$(BINARY_NAME)"

test:
	@echo "Running tests..."
	$(TEST) -v

run:
	$(GO) run $(MAIN_FILE)

clean:
	@echo "Cleaning..."
	$(CLEAN)
	rm -f $(BINARY_NAME) $(MOD_FILE)
	@echo "Clean complete"

.PHONY: all init build test clean run