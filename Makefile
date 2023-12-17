.PHONY: build
build:
	go build -v ./cmd/main/main.go
run-server:
	sudo docker-compose up -d
check-docker-process:
	sudo docker-compose ps

.DEFAULT_GOAL := build

