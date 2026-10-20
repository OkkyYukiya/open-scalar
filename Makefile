# Common development commands. Toolchain is managed by mise, releases by GoReleaser.

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the binary into ./openscalar
	go build -o openscalar ./cmd/openscalar

.PHONY: run
run: ## Run the TUI
	go run ./cmd/openscalar

.PHONY: test
test: ## Run tests with the race detector
	go test -race ./...

.PHONY: cover
cover: ## Run tests and open the HTML coverage report
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

.PHONY: fmt
fmt: ## Format all Go files
	gofmt -w .

.PHONY: lint
lint: ## Run gofmt check, go vet and staticcheck (same as CI)
	@out=$$(gofmt -l .); if [ -n "$$out" ]; then echo "not gofmt-ed:"; echo "$$out"; exit 1; fi
	go vet ./...
	staticcheck ./...

.PHONY: tidy
tidy: ## Tidy go.mod and go.sum
	go mod tidy

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf openscalar dist cover.out

.PHONY: release
release: lint test ## Tag and push a release: make release tag=v0.1.0
	@if [ -z "$(tag)" ]; then echo "usage: make release tag=v0.1.0"; exit 1; fi
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'
