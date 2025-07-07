GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=shepherd
DOCKER_REPOSITORY_OWNER=alwindoss
VERSION=0.0.1

.PHONY: build
build: clean generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ ./cmd/shepherd/...

.PHONY: clean
clean:
	rm -rf bin
	go clean

.PHONY: run
run: build
	./bin/shepherd

.PHONY: dev
dev:
	go tool air

.PHONY: setup
setup:
	go mod tidy
	go get -v ./...

.PHONY: generate
generate:
	go tool templ generate

.PHONY: package
package:
	docker build -t $(DOCKER_REPOSITORY_OWNER)/$(BINARY_NAME):$(VERSION) .

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/ -v ./cmd/$(BINARY_NAME)/...