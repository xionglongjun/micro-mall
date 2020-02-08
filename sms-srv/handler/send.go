package handler

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/xionglongjun/micro-mall/sms-srv/enums"
	"github.com/xionglongjun/micro-mall/sms-srv/models"
	"github.com/xionglongjun/micro-mall/sms-srv/modules/provider"
	send "github.com/xionglongjun/micro-mall/sms-srv/proto/send"
	"github.com/xionglongjun/micro-mall/sms-srv/repository"
	"github.com/xionglongjun/micro-mall/sms-srv/utils"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
)

// Send ...
type Send struct {
	Repo        *repository.SendRepo
	SmsProvider provider.Driver
}

// BizType is a single request handler called via client.Call or the generated client code
func (e *Send) BizType(ctx context.Context, req *send.Request, rsp *send.BizTypeResponse) error {
	var typeList []*send.MapType
	for key, val := range send.BizType_name {
		fmt.Println(key, val)
		typeList = append(typeList, &send.MapType{
			Key:   key,
			Value: val,
		})
	}
	rsp.List = typeList
	return nil
}

// Code is a single request handler called via client.Call or the generated client code
func (e *Send) Code(ctx context.Context, req *send.CodeRequest, rsp *send.CodeResponse) error {
	var (
		err           error
		content       string
		sid           string
		code          string
		expire        time.Duration
		expiresAt     time.Time
		now           = time.Now()
		re            *regexp.Regexp
		repo          = e.Repo
		templateRepo  = repository.TemplateRepo{DB: repo.DB}
		codeRepo      = repository.CodeRepo{DB: repo.DB}
		smsProvider   = e.SmsProvider
		templateModel *models.Template
	)

	if req.Expires <= 0 {
		expire = 5 * time.Minute
	}
	expiresAt = now.Add(expire)
	if !utils.ValidateMobile(req.Mobile) {
		return errors.BadRequest("go.micro.srv.sms Code", "mobile format fail")
	}
	templateModel, err = templateRepo.GetCodeTemplateFirst(smsProvider.String(), uint(req.BizType))
	if err != nil {
		return errors.BadRequest("go.micro.srv.sms Code", err.Error())
	}
	code = utils.GenValidateCode(6)
	re = regexp.MustCompile("#([a-z|A-Z]*)#")
	content = re.ReplaceAllString(templateModel.Content, code)
	sendModel := &models.Send{
		Content:  content,
		Mobile:   req.Mobile,
		Provider: smsProvider.String(),
		BizType:  uint(req.BizType),
	}
	if err = e.Repo.Create(sendModel); err != nil {
		return errors.InternalServerError("go.micro.srv.sms Code", err.Error())
	}
	sid, err = smsProvider.Send(req.Mobile, content)
	if err != nil {
		e.Repo.Model(sendModel).Update("message", err.Error())
		return errors.BadRequest("go.micro.srv.sms Code", "%s, %s", err.Error())
	}
	if err := e.Repo.Model(sendModel).Update(&models.Send{
		SID:     sid,
		Success: uint(enums.Yes),
	}).Error; err != nil {
		log.Infof("save send status: %s", err.Error())
	} else {
		codeRepo.Create(&models.Code{
			BizType:   uint(req.BizType),
			Code:      code,
			Mobile:    req.Mobile,
			Send:      *sendModel,
			ExpiresAt: expiresAt,
		})
	}

	rsp.ExpiresAt = &timestamp.Timestamp{
		Seconds: expiresAt.Unix(),
	}
	rsp.Success = true
	return nil
}

// Validate is a server side stream handler called via client.Stream or the generated client code
func (e *Send) Validate(ctx context.Context, req *send.ValidateRequest, rsp *send.ValidateResponse) error {
	var (
		err      error
		repo     = e.Repo
		codeRepo = repository.CodeRepo{DB: repo.DB}
	)
	if ok, _ := regexp.MatchString(`([0-9]\d{5})$`, req.Code); !ok {
		return errors.BadRequest("go.micro.srv.sms Validate", "code faild")
	}

	if !utils.ValidateMobile(req.Mobile) {
		return errors.BadRequest("go.micro.srv.sms Validate", "mobile format fail")
	}
	if _, err = codeRepo.Validate(req.Mobile, req.Code, uint(req.BizType)); err != nil {
		return errors.BadRequest("go.micro.srv.sms Validate", "%s", err.Error())
	}

	rsp.Success = true
	return nil
}
