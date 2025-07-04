.PHONY: build
build: clean
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
	go tool air

.PHONY: setup
setup:
	go mod tidy