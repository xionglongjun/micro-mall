package handler

import (
	"context"
	"sms-srv/modules/provider"
	send "sms-srv/proto/send"
	"sms-srv/repository"

	"github.com/micro/go-micro/errors"
)

// Send ...
type Send struct {
	Repo         *repository.SendRepo
	CodeRepo     *repository.CodeRepo
	TemplateRepo *repository.TemplateRepo
	SmsProvider  provider.Driver
}

// Code is a single request handler called via client.Call or the generated client code
func (e *Send) Code(ctx context.Context, req *send.CodeRequest, rsp *send.CodeResponse) error {
	if req.Mobile == "" {
		return errors.BadRequest("go.micro.srv.sms Code", "mobile not null")
	}

	e.TemplateRepo.GetCodeTemplateFirst("yunpian", 0)
	_, err := e.SmsProvider.Send(req.Mobile, req.Code)
	if err != nil {
		return errors.BadRequest("go.micro.srv.sms Code", err.Error())
	}
	rsp.Success = true
	return nil
}

// Validate is a server side stream handler called via client.Stream or the generated client code
func (e *Send) Validate(ctx context.Context, req *send.ValidateRequest, rsp *send.ValidateResponse) error {

	return nil
}
