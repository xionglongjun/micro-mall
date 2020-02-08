module github.com/xionglongjun/micro-mall/user-srv

go 1.13

// replace github.com/xionglongjun/micro-mall/sms-srv => /Users/jayden/Development/Code/Golang/micro/micro-mall/sms-srv

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/xionglongjun/micro-mall/sms-srv v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
)
