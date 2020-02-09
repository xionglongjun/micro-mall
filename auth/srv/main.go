package main

import (
	"auth-srv/handler"
	"auth-srv/repository"
	"auth-srv/subscriber"
	"auth-srv/utils"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	auth "auth-srv/proto/auth"
)

func main() {

	jwt := utils.Jwt{
		Secret:  "ldjkvndmdhworywwnxksdhsdksdf",
		Issuer:  "user",
		Expires: 24 * time.Hour,
	}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{Repo: &repository.AuthRepo{DB: db}, Jwt: jwt})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.auth", service.Server(), new(subscriber.Auth))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.auth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
