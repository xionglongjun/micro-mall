GOPATH:=$(shell go env GOPATH)

.PHONY: proto
type = srv
mod = user
name = user
proto:
	@echo execute $(mod) proto file generate
	protoc --proto_path=. --micro_out=. --go_out=. proto/$(type)/$(mod)/${name}.proto
	
.PHONY: api
api:
	go run main.go --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats api --handler=api --namespace=go.mall 

.PHONY: user-srv
user-srv:
	go run user-srv/main.go
.PHONY: web
web:
	go run web/main.go web

.PHONY: user-api
user-api:
	go run user-api/main.go

.PHONY: build
build: proto

	go build -o user-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-api:latest