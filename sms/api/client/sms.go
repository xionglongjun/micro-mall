package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	sms "github.com/xionglongjun/micro-mall/sms-api/proto/sms"
)

type smsKey struct {}

// FromContext retrieves the client from the Context
func SmsFromContext(ctx context.Context) (sms.SmsService, bool) {
	c, ok := ctx.Value(smsKey{}).(sms.SmsService)
	return c, ok
}

// Client returns a wrapper for the SmsClient
func SmsWrapper(service micro.Service) server.HandlerWrapper {
	client := sms.NewSmsService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, smsKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
