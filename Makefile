MODULE_NAME := $(shell go list -m)
GOCMD ?= go
PACKAGES := ./...

.PHONY: all
all: build

.PHONY: build
build:
	@echo "==> Building ${MODULE_NAME}..."
	GOOS="linux" GOARCH="amd64" goreleaser build --snapshot --single-target --clean

.PHONY: clean
clean:
	@echo "==> Cleaning..."
	rm -rf bin dist _output

.PHONY: fmt
fmt:
	@echo "==> Formatting source..."
	$(GOCMD) fmt $(PACKAGES)

.PHONY: test
test:
	@echo "==> Run tests..."
	$(GOCMD) test ./...

.PHONY: tidy
tidy:
	@echo "==> Tidying module_NAMEs..."
	$(GOCMD) mod tidy

.PHONY: lint
lint:
	@echo "==> Linting source..."
	@if command -v golangci-lint > /dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed; run 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest'"; \
		exit 1; \
	fi

.PHONY: deps
deps:
	@echo "==> Downloading dependencies..."
	$(GOCMD) mod download

.PHONY: release
release:
	@echo "==> Building release binaries..."
	goreleaser release --snapshot --clean
