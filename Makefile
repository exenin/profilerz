# Variables
APP_NAME = profilerz
GO_BUILD_CMD = go build -o $(APP_NAME) ./cmd

# Default target
.PHONY: all
all: build

# Build the Go binary
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	$(GO_BUILD_CMD)

# Run the Go program (you can specify subcommands)
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME) $(ARGS)

# Test the Go code
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Clean up binary and other build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
	go clean

# Run Go fmt to format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run Go vet to check for common issues
.PHONY: vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run Go lint (requires golangci-lint installed)
.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Tidy up dependencies in go.mod
.PHONY: tidy
tidy:
	@echo "Tidying up go.mod and go.sum..."
	go mod tidy

# Install the binary globally
.PHONY: install
install:
	@echo "Installing $(APP_NAME)..."
	go install ./cmd

# Uninstall the binary
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(APP_NAME)..."
	rm -f $(GOPATH)/bin/$(APP_NAME)

# Build and run the application (e.g., with ARGS="init" or ARGS="profile")
.PHONY: build-run
build-run: build run
