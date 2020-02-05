
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
mod = sms
name = sms
proto:
	@echo execute $(mod) proto file generate
	protoc --proto_path=. --micro_out=. --go_out=. proto/$(mod)/${name}.proto
	
.PHONY: build
build: proto

	go build -o sms-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t sms-srv:latest
