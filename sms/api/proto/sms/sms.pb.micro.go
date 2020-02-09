// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sms/sms.proto

package go_micro_api_sms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Sms service

type SmsService interface {
	Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type smsService struct {
	c    client.Client
	name string
}

func NewSmsService(name string, c client.Client) SmsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.sms"
	}
	return &smsService{
		c:    c,
		name: name,
	}
}

func (c *smsService) Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Sms.Call", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sms service

type SmsHandler interface {
	Call(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterSmsHandler(s server.Server, hdlr SmsHandler, opts ...server.HandlerOption) error {
	type sms interface {
		Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Sms struct {
		sms
	}
	h := &smsHandler{hdlr}
	return s.Handle(s.NewHandler(&Sms{h}, opts...))
}

type smsHandler struct {
	SmsHandler
}

func (h *smsHandler) Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SmsHandler.Call(ctx, in, out)
}