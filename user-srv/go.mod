module user-srv

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/golang/protobuf v1.3.3
	github.com/google/btree v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.27 // indirect
	github.com/nats-io/nats-server/v2 v2.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/xionglongjun/micro-mall/sms-srv v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/tools v0.0.0-20191216173652-a0e659d51361 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1 // indirect
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

// replace github.com/xionglongjun/micro-mall => /Users/jayden/Development/Code/Golang/micro-mall
replace github.com/xionglongjun/micro-mall/sms-srv => /Users/jayden/Development/Code/Golang/micro/micro-mall/sms-srv
