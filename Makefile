.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./main.go ./cmd ./internal

format:
	@pnpm run format && gofmt -w ./main.go ./cmd ./internal

test:
	go test -race -cover -coverprofile=coverage.txt ./cmd/... ./internal/...

build:
	go build -race -mod vendor -o spectacle
