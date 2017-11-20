.PHONY: default build install

default: build

build:
	go build -o gen .
	mv gen ${GOPATH}/bin/