package handler

import (
	auth "auth-srv/proto/auth"
	"auth-srv/repository"
	"auth-srv/utils"
	"context"
	"user-srv/models"

	send "github.com/xionglongjun/micro-mall/sms-srv/proto/send"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
)

// Auth ...
type Auth struct {
	Repo *repository.AuthRepo
	Jwt  utils.Jwt
}

// Name ...
func (h *Auth) Name(ctx context.Context, req *auth.NameRequest, rsp *auth.AuthResponse) error {
	var (
		err       error
		userModel *models.User
	)
	if req.Name == "" {
		return errors.BadRequest("go.micro.srv.user Name", "name not null")
	}

	if req.Password == "" {
		return errors.BadRequest("go.micro.srv.user Name", "Password not null")
	}

	userModel, err = h.Repo.Name(req.Name)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Name", "name find null")
	}

	pass, err := utils.EncodeSalt("123456", userModel.Salt)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Name", "Password encode validate fail")
	}

	if pass != userModel.Password {
		return errors.BadRequest("go.micro.srv.user Name", "Password fail")
	}
	claims := utils.MyClaims{
		ID:     userModel.ID,
		Name:   userModel.Name,
		Mobile: userModel.Mobile,
	}
	jwtToken, err := h.Jwt.NewToken(claims)
	if err != nil {
		return errors.Unauthorized("go.micro.srv.user Name", "token generate fail")
	}

	rsp.Token = jwtToken.Token
	rsp.ExpiresAt = jwtToken.ExpiresAt.Format("2006-01-02 15:04:05")
	return nil
}

// Mobile ...
func (h *Auth) Mobile(ctx context.Context, req *auth.MobileRequest, rsp *auth.AuthResponse) error {
	var (
		err       error
		userModel *models.User
	)
	smsClient := send.NewSendService("go.micro.srv.sms", client.DefaultClient)
	if !utils.ValidateMobile(req.Mobile) {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile not null")
	}

	if req.Code == "" {
		return errors.BadRequest("go.micro.srv.user Mobile", "code not null")
	}
	_, err = smsClient.Validate(ctx, &send.ValidateRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "sms validate: %s", errors.Parse(err.Error()).Detail)
	}
	userModel, err = h.Repo.Mobile(req.Mobile)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile find null")
	}

	claims := utils.MyClaims{
		ID:     userModel.ID,
		Name:   userModel.Name,
		Mobile: userModel.Mobile,
	}
	jwtToken, err := h.Jwt.NewToken(claims)
	if err != nil {
		return errors.Unauthorized("go.micro.srv.user Name", "token generate fail")
	}
	rsp.Token = jwtToken.Token
	rsp.ExpiresAt = jwtToken.ExpiresAt.Format("2006-01-02 15:04:05")
	return nil
}
