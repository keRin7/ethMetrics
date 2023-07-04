
CGO_ENABLED=0

.PHONY: build
build: 
	go build -o bin/app -v ./cmd
#	go build -a -installsuffix cgo -o bin/app -v ./cmd

.DEFAULT_GOAL := build