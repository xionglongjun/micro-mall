GOPATH:=$(shell go env GOPATH)
NETWORK = micro-mall

api:
	docker run --rm --name micro-mall-api -d -p 8080:8080 --network ${NETWORK} -e MICRO_REGISTER=mdns micro/micro api --handler=rpc --address=:8080

web:
	docker run --rm --name micro-mall-web -d -p 8082:8082 --network ${NETWORK} micro/micro web  --address=:8082

.PHONY: api-self
api-self:
	go run main.go --cors-allowed-headers="Content-Type,Authorization,Client-Type" --cors-allowed-origins="*" --cors-allowed-methods="OPTIONS,DELETE,GET,POST,PUT" --enable_stats api --handler=api --namespace=go.mall 

PHONY: api web api-self