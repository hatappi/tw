export GOPATH := $(shell go env GOPATH)
export GOBIN := $(GOPATH)/bin
export PATH := $(GOBIN):${shell pwd}/bin:$(PATH)
export GIT_HASH := `git rev-parse --short HEAD`

GOBIN:=${PWD}/bin
PATH:=${GOBIN}:${PATH}

install-tools:
	./scripts/install_tools.sh

build:
	go build \
	  -ldflags "-X github.com/hatappi/tw/cmd.version=commit-${GIT_HASH} -X github.com/hatappi/tw/cmd.commit=${GIT_HASH}" \
	  -o ./dist/tw \
	  main.go

.PHONY: lint
lint:
	${GOBIN}/golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	${GOBIN}/golangci-lint run --fix ./...

.PHONY: dependencies
dependencies:
	@GO111MODULE=off go get -u github.com/Songmu/ghch/cmd/ghch
	@go mod download
	@go mod tidy

.PHONY: test
test:
	@go test ./...
