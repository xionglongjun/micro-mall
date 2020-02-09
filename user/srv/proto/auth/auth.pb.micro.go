// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Client API for Auth service

type AuthService interface {
	Name(ctx context.Context, in *NameRequest, opts ...client.CallOption) (*AuthResponse, error)
	Mobile(ctx context.Context, in *MobileRequest, opts ...client.CallOption) (*AuthResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Name(ctx context.Context, in *NameRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Name", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Mobile(ctx context.Context, in *MobileRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Mobile", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Name(context.Context, *NameRequest, *AuthResponse) error
	Mobile(context.Context, *MobileRequest, *AuthResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Name(ctx context.Context, in *NameRequest, out *AuthResponse) error
		Mobile(ctx context.Context, in *MobileRequest, out *AuthResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		ValidateToken(ctx context.Context, in *Token, out *Token) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Name(ctx context.Context, in *NameRequest, out *AuthResponse) error {
	return h.AuthHandler.Name(ctx, in, out)
}

func (h *authHandler) Mobile(ctx context.Context, in *MobileRequest, out *AuthResponse) error {
	return h.AuthHandler.Mobile(ctx, in, out)
}

func (h *authHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AuthHandler.Register(ctx, in, out)
}

func (h *authHandler) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}