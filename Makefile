PKGS := $(shell go list ./...)

install-deps:
	go get

.PHONY: test
test:
	go test $(PKGS)

build: install-deps test
	go build cmd/server/main.go

start-server:
	go run cmd/server/main.go