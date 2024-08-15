#include .env.example
#export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

.PHONY: build-fluent
build-fluent:
	# @go build -ldflags "-X main.version=$(shell git describe --abbrev=0 --tags)" -o fluent
	go build cmd/fluent.go


PHONY: test
test: ## run go tests
	go test -v ./...

PHONY: PB
PB: ## build pocketbase
	go build app/pb/pbserver.go

integration-test: ### run integration-test
	go clean -testcache && go test -v ./integration-test/...
.PHONY: integration-test


.PHONY: build-cli
build-cli:
	go build cmd/fluent.go

.PHONY: format
format:
	gofmt -l -w .

.PHONY: fmtcheck
fmtcheck: ## run gofmt and print detected files
	@sh -c "'$(CURDIR)/ci/scripts/goformat.sh'"