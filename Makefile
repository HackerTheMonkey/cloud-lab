all: install test

install:
	@go install

test:
	@go test ./... --count=1 -v

run:
	@codematters
