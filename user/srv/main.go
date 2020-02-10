package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xionglongjun/micro-mall/user/srv/handler"
	"github.com/xionglongjun/micro-mall/user/srv/models"
	auth "github.com/xionglongjun/micro-mall/user/srv/proto/auth"
	user "github.com/xionglongjun/micro-mall/user/srv/proto/user"

	"github.com/xionglongjun/micro-mall/user/srv/repository"
	"github.com/xionglongjun/micro-mall/user/srv/utils"

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
	var dbHost = os.Getenv("DB_HOST")
	var dbName = os.Getenv("DB_NAME")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	fmt.Println(dbHost)
	if dbHost == "" {
		dbHost = "127.0.0.1:3306"
	}
	if dbName == "" {
		dbName = "mall_user"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbPassword == "" {
		dbPassword = "123456"
	}
	mysql := &models.Mysql{
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		Host:     dbHost,
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
