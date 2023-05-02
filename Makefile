#!make
#----------------------------------------
# Settings
#----------------------------------------
.DEFAULT_GOAL := help

#--------------------------------------------------
# Variables
#--------------------------------------------------
TEST?=$$(go list ./...)
GO_FILES?=$$(find . -name '*.go')

#--------------------------------------------------
# Targets
#--------------------------------------------------
.PHONY: bootstrap
bootstrap: ## Downloads and cleans up all dependencies
	@go mod tidy
	@go mod download

.PHONY: fmt
fmt: ## Formats go files
	@echo "==> Formatting files..."
	@gofmt -w -s $(GO_FILES)
	@echo ""

.PHONY: check
check: ## Checks code for linting/construct errors
	@echo "==> Checking if files are well formatted..."
	@gofmt -l $(GO_FILES)
	@echo ""
	@echo "==> Checking if files pass go vet..."
	@go list -f '{{.Dir}}' ./... | xargs go vet;
	@echo ""

.PHONY: test
test: check ## Runs all tests
	@echo "==> Running tests..."
	@go test -v --race $(TEST) -parallel=20
	@echo ""

.PHONY: coverage
coverage: ## Runs code coverage
	@mkdir -p .coverage
	@go test --p=1 $(TEST) -coverprofile=.coverage/cover.out -covermode=atomic
	@go tool cover -html=.coverage/cover.out

.PHONY: package
package: clean ## Packages the lambda code into a ZIP file
	@echo "==> Packaging application ..."
	@mkdir -p .dist
	@rsync -a . .dist --exclude .dist --exclude .gitignore --exclude Makefile --exclude README.md
	@cd .dist && zip -qr ../dist.zip . && cd ..
	@echo ""

clean: ## Cleans up temporary and compiled files
	@echo "==> Cleaning up ..."
	@rm -rf .coverage
	@rm -rf .dist
	@rm -rf dist.zip
	@echo ""

help:
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
