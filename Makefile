#--------------------------------------------------
# Variables
#--------------------------------------------------
SCRIPT=$(CURDIR)/scripts
#--------------------------------------------------
GOFMT_FILES?=$$(find . -name '*.go')
#--------------------------------------------------
PACKAGES?=$$(go list ./... \
	| grep -vF function/internal )

#--------------------------------------------------
# Targets
#--------------------------------------------------
package: clean fmt
	@mkdir -p dist
	@rsync -a . dist \
		--exclude dist \
		--exclude scripts \
		--exclude Makefile \
		--exclude .gitignore \
		--exclude *_test.go
	@cd dist && GOOS=linux go build -o main
	@cd dist && zip -qr ../lambda.zip .
	@rm -rf dist

fmt:
	@gofmt -w $(GOFMT_FILES)

test:
	@go test $(PACKAGES) -parallel=20

coverage:
	@sh -c "'$(SCRIPT)/getCoverage.sh'"

coverage-html:
	@SHOW_HTML='true' sh -c "'$(SCRIPT)/getCoverage.sh'"

.PHONY: clean fmt package test
clean:
	@rm -rf dist
	@rm -f lambda.zip