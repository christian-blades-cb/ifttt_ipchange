GOARCH ?= arm
GOOS ?= linux
GO15VENDOREXPERIMENT := 1
BUILDTAG ?= dev
.DEFAULT_GOAL := dockerize

.phony: dockerize

build:
	go build

dockerize: build
	docker build -t christianbladescb/ifttt_ipchange:${GOARCH}-${GOOS}-${BUILDTAG} .

run: build
	go build .
