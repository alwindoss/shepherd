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
	go tool air

.PHONY: setup
setup:
	go mod tidy

.PHONY: copy
copy: build
	cp ./bin/shepherd /home/alwin/Sandbox/opt/go/bin/