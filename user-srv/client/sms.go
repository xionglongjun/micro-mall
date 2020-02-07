package client

import (
	"context"
	"fmt"

	send "github.com/xionglongjun/micro-mall/sms-srv/proto/send"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// SmsClient ...
type SmsClient struct {
	Client client.Client
}

// ValidateCode ...
func (s *SmsClient) ValidateCode(mobile, code string) {
	var (
		client = s.Client
	)
	req := client.NewRequest("go.micro.srv.sms", "Send.Validate", send.ValidateRequest{
		Mobile: mobile,
		Code:   code,
	})
	fmt.Println(req)
}

type smskey struct{}

// SmsContext retrieves the client from the Context
func SmsContext(ctx context.Context) (send.SendService, bool) {
	c, ok := ctx.Value(smskey{}).(send.SendService)
	return c, ok
}

// SmsWrapper returns a wrapper for the UserClient
func SmsWrapper(service micro.Service) server.HandlerWrapper {
	client := send.NewSendService("go.micro.srv.sms", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, smskey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
