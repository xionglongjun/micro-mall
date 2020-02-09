
GOPATH:=$(shell go env GOPATH)



.PHONY: proto
name = auth
proto:
	@echo execute ${name} proto file generate
	protoc --proto_path=. --micro_out=. --go_out=. proto/${name}/${name}.proto
	
.PHONY: build
build: proto

	go build -o auth-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t auth-srv:latest
