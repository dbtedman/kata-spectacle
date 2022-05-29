.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./cmd ./internal

format:
	@pnpm run format && gofmt -w ./cmd ./internal

test:
	go test -race -cover -coverprofile=coverage.txt ./cmd/... ./internal/...

build:
	go build -race -mod vendor -o spectacle ./cmd/spectacle
