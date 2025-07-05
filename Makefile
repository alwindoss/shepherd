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
build: clean
	cd ui && npm run build
	go build -o bin/ ./cmd/shepherd/...

.PHONY: clean
clean:
	rm -rf bin
	go clean

.PHONY: run
run: build
	./bin/shepherd

.PHONY: dev
dev:
	go tool task --parallel backend frontend

.PHONY: setup
setup:
	cd ui && npm i
	go mod tidy
	go get -v ./...

.PHONY: copy
copy: build
	@echo "Copying to:" ${GOBIN}
	cp ./bin/shepherd ${GOBIN}

.PHONY: package
package:
	docker build -t $(DOCKER_REPOSITORY_OWNER)/$(BINARY_NAME):$(VERSION) .