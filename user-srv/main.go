package main

import (
	"context"
	"fmt"
	"time"
	"user-srv/client"
	"user-srv/handler"
	"user-srv/models"
	"user-srv/repository"
	"user-srv/utils"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

	auth "user-srv/proto/auth"
	user "user-srv/proto/user"
)

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// log.Fatalf("[wrapper] server request: %v", req.Endpoint())
		fmt.Printf("[wrapper] [%v] server request: %s", time.Now(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func main() {
	var (
		driver models.Driver
		jwt    utils.Jwt
	)
	driver = &models.Mysql{
		User:     "root",
		Password: "123456",
		Name:     "mall_user",
		Host:     "127.0.0.1:3306",
		Debug:    true,
	}

	jwt.Secret = "ldjkvndmdhworywwnxksdhsdksdf"
	jwt.Issuer = "user"
	jwt.Expires = 24 * time.Hour

	db, err := driver.Connection()
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	// 数据迁移
	db.Begin().AutoMigrate(&models.User{})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(logWrapper),
	)

	// Initialise service
	service.Init(
		micro.WrapHandler(client.SmsWrapper(service)),
	)

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{Repo: &repository.AuthRepo{DB: db}, Jwt: jwt})
	user.RegisterUserHandler(service.Server(), &handler.User{Repo: &repository.UserRepo{DB: db}})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
