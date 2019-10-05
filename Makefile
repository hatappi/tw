build:
	go build -o ./dist/tw main.go

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix ./...

