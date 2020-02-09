GOPATH:=$(shell go env GOPATH)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
# IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
IMAGE_TAG = v0.0.1
NAME = user-srv
IMAGE_NAME = micro-mall/${NAME}
NETWORK = micro-mall
MYSQL_HOST = mysql:3306
PROTO = user

all: build

proto:
	@echo execute ${PROTO} proto file generate
	protoc --proto_path=. --micro_out=. --go_out=. proto/${PROTO}/${PROTO}.proto
	
vendor:
	go mod vendor

build:
	go build -o ${NAME} *.go

test:
	go test -v ./... -cover

docker:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):latest

run:
	docker run --rm --name ${NAME} -p :50051 --network ${NETWORK} -e MICRO_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e MYSQL_HOST=${MYSQL_HOST} ${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: build proto clean vet test docker run