export GOPATH := $(shell go env GOPATH)
export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
export GIT_HASH := `git rev-parse --short HEAD`

build:
	go build \
	  -ldflags "-X github.com/hatappi/tw/cmd.version=commit-${GIT_HASH} -X github.com/hatappi/tw/cmd.commit=${GIT_HASH}" \
	  -o ./dist/tw \
	  main.go

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix ./...

lint-dependencies:
	@GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: dependencies
dependencies:
	@go mod download

.PHONY: test
test:
	@go test ./...
