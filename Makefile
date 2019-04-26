.PHONY: get test

get:
	dep ensure

build:
	go build -o kubeconfig-factory

test:
	go test -v ./...
