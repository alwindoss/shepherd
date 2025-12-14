.PHONY: build run clean

build:
	@echo "Building the project..."
	@echo "Building UI..."
	cd ui && npm ci && npm run build

	@echo "Building Go binary..."
	go build -o bin/shepherd

run: build
	@echo "Running the project..."
	./bin/shepherd

clean:
	@echo "Cleaning up..."
	rm -rf bin/