.PHONY: build run clean

build:
	@echo "Building the project..."
	go build -o bin/

run: build
	@echo "Running the project..."
	./bin/shepherd

clean:
	@echo "Cleaning up..."
	rm -rf bin/