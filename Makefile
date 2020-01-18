.PHONY: install test 

install:
	go get

test: install
	go test ./...
