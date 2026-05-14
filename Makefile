.PHONY: setup build build-ui generate run dev build-backend

setup:
	go tool task setup

build:
	go tool task build
	
build-backend: generate
	go tool task build-backend
	
build-ui:
	go tool task build-ui

generate:
	go tool task generate

run: build-ui generate build-backend
	go tool task run

dev: setup
	go tool task dev

container-build:
	go tool task container-build

container-run:
	go tool task container-run