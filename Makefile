
BINARY = ./bin/reverse-proxy

.PHONY: clean
clean:
	@rm -f $(BINARY)

.PHONY: build
build: clean
	@go build -o $(BINARY) ./cmd/main.go

.PHONY: docker-build
docker-build:
	@docker build -t ghcr.io/trent-j/reverse-proxy:v2 -f Dockerfile .
