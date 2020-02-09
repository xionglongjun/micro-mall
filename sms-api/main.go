package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/xionglongjun/micro-mall/sms-api/client"
	"github.com/xionglongjun/micro-mall/sms-api/handler"

	sms "github.com/xionglongjun/micro-mall/sms-api/proto/sms"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.sms"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Sms srv client
		micro.WrapHandler(client.SmsWrapper(service)),
	)

	// Register Handler
	sms.RegisterSmsHandler(service.Server(), new(handler.Sms))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
