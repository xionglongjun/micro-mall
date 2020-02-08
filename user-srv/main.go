package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-srv/handler"
	"user-srv/models"
	auth "user-srv/proto/auth"
	user "user-srv/proto/user"
	"user-srv/repository"
	"user-srv/utils"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
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
	mysql := &models.Mysql{
		User:     "root",
		Password: "123456",
		Name:     "mall_user",
		Host:     "127.0.0.1:3306",
		Debug:    true,
	}
	jwt := utils.Jwt{
		Secret:  "ldjkvndmdhworywwnxksdhsdksdf",
		Issuer:  "user",
		Expires: 24 * time.Hour,
	}

	db, err := mysql.Connection()
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
		// 日志包装器
		micro.WrapHandler(logWrapper),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{Repo: &repository.AuthRepo{DB: db}, Jwt: jwt})
	user.RegisterUserHandler(service.Server(), &handler.User{Repo: &repository.UserRepo{DB: db}})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}