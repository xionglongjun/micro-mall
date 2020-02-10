GOPATH:=$(shell go env GOPATH)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
# IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
IMAGE_TAG = v0.0.1
NAME = micro-api
IMAGE_NAME = micro-mall/${NAME}
NETWORK = micro-mall
MICRO_REGISTER = mdns
TYPE = api
HANDLER = api

api:
	go run main.go --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats ${TYPE} --handler=${HANDLER} --namespace=go.mall 


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
	docker run --rm --name ${NAME} -d -p 8080:8080 --network ${NETWORK} -e MICRO_REGISTER=${MICRO_REGISTER} ${IMAGE_NAME}:${IMAGE_TAG} ${TYPE} --handler=${HANDLER} --address=:8080


PHONY: docker run build
