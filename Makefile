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