.PHONY: test 

install:
	go get

test: install
	go test ./...
