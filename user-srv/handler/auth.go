package handler

import (
	"context"
	"github.com/xionglongjun/micro-mall/user-srv/models"
	auth "github.com/xionglongjun/micro-mall/user-srv/proto/auth"
	"github.com/xionglongjun/micro-mall/user-srv/repository"
	"github.com/xionglongjun/micro-mall/user-srv/utils"

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

	pass, err := utils.EncodeSalt(req.Password, userModel.Salt)
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

	if !utils.ValidateMobile(req.Mobile) {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile not null")
	}

	if req.Code == "" {
		return errors.BadRequest("go.micro.srv.user Mobile", "code not null")
	}
	
	userModel, err = h.Repo.Mobile(req.Mobile)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile find null")
	}

	smsClient := send.NewSendService("go.micro.srv.sms", client.DefaultClient)
	_, err = smsClient.Validate(ctx, &send.ValidateRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
	})

	if err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "sms validate: %s", errors.Parse(err.Error()).Detail)
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

// Register ...
func (h *Auth) Register(ctx context.Context, req *auth.RegisterRequest, rsp *auth.RegisterResponse) error {
	var (
		err       error
		userRepo = repository.UserRepo{DB: h.Repo.DB}
	)
	salt := utils.EncodeMD5(utils.GenValidateCode(10))
	pass, err := utils.EncodeSalt(req.Password, salt)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Name", "Password encode validate fail")
	}

	if !utils.ValidateMobile(req.Mobile) {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile not null")
	}

	if req.Code == "" {
		return errors.BadRequest("go.micro.srv.user Mobile", "code not null")
	}
	
	_, err = h.Repo.Mobile(req.Mobile)
	if err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "mobile find null")
	}
	
	smsClient := send.NewSendService("go.micro.srv.sms", client.DefaultClient)
	_, err = smsClient.Validate(ctx, &send.ValidateRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
		BizType: send.BizType_register,
	})
	userModel := models.User{
		Mobile: req.Mobile,
		Name: req.Name,
		Password: pass,
		Salt: salt,
	}
	if err = userRepo.Create(&userModel); err != nil {
		return errors.BadRequest("go.micro.srv.user Mobile", "register fail")
	}
	rsp.Success = true
	return nil
}

// ValidateToken ...
func (h *Auth) ValidateToken(ctx context.Context, req *auth.Token, rsp *auth.Token) error {
	if req.Token == "" {
		return errors.BadRequest("go.micro.srv.user ValidateToken", "token not null")
	}
	myClaims, err := h.Jwt.ParseToken(req.Token)
	if err != nil {
		return errors.Unauthorized("go.micro.srv.user ValidateToken", err.Error())
	}
	if myClaims.ID == 0 {
		return errors.Unauthorized("go.micro.srv.user ValidateToken", "invalid user")
	}
	rsp.Valid = true
	return nil
}