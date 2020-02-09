
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
name = sms
proto:
	@echo execute ${name} proto file generate
	protoc --proto_path=${GOPATH}/src:.  --micro_out=. --go_out=. proto/${name}/${name}.proto
	
.PHONY: build
build: proto

	go build -o sms-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t sms-api:latest
