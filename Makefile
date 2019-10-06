build:
	go build -o ./dist/tw main.go

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix ./...

lint-dependencies:
	@GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: dependencies
dependencies:
	@go mod download

.PHONY: test
test:
	@go test ./...
