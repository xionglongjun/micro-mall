package main

import (
	"fmt"
	"os"

	"github.com/xionglongjun/micro-mall/sms-srv/handler"
	"github.com/xionglongjun/micro-mall/sms-srv/models"
	"github.com/xionglongjun/micro-mall/sms-srv/modules/provider"
	"github.com/xionglongjun/micro-mall/sms-srv/repository"
	"github.com/xionglongjun/micro-mall/sms-srv/subscriber"

	send "github.com/xionglongjun/micro-mall/sms-srv/proto/send"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	var host = os.Getenv("MYSQL_HOST")
	var name = os.Getenv("MYSQL_NAME")
	var user = os.Getenv("MYSQL_USER")
	var password = os.Getenv("MYSQL_PASSWORD")
	fmt.Println(host)
	if host == "" {
		host = "127.0.0.1:3306"
	}
	if name == "" {
		name = "mall_sms"
	}
	if user == "" {
		user = "root"
	}
	if password == "" {
		password = "123456"
	}
	mysql := &models.Mysql{
		User:     user,
		Password: password,
		Name:     name,
		Host:     host,
		Debug:    true,
	}

	db, err := mysql.Connection()
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	// 数据迁移
	db.Begin().AutoMigrate(&models.Send{}, &models.Code{}, &models.Template{})
	// APIkey yunpian sms apikey
	APIKey := "7aa71cd46746dfec1c69a3c62f443722"

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.sms"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	send.RegisterSendHandler(service.Server(), &handler.Send{
		Repo: &repository.SendRepo{DB: db},
		SmsProvider: &provider.YunPian{
			APIKey: APIKey,
			Debug:  true,
		},
	})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.sms", service.Server(), new(subscriber.Send))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
