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
	go mod tidy
	go get -v ./...

.PHONY: copy
copy: build
	@echo "Copying to:" ${GOBIN}
	cp ./bin/shepherd ${GOBIN}

.PHONY: build-ui
build-ui:
	cd ui && npm i

.PHONY: package
package:
	docker build -t $(DOCKER_REPOSITORY_OWNER)/$(BINARY_NAME):$(VERSION) .

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/ -v ./cmd/$(BINARY_NAME)/...