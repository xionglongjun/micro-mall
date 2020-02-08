package handler

import (
	"context"
	"github.com/xionglongjun/micro-mall/user-srv/models"
	user "github.com/xionglongjun/micro-mall/user-srv/proto/user"
	"github.com/xionglongjun/micro-mall/user-srv/repository"

	"github.com/micro/go-micro/errors"
)

// User ...
type User struct {
	Repo *repository.UserRepo
}

// Create ...
func (h *User) Create(ctx context.Context, req *user.CreateRequest, rsp *user.CreateResponse) error {
	if err := h.Repo.Create(&models.User{
		Name: req.Name,
		Mobile: req.Mobile,
	}); err != nil {
		return errors.New("go.micro.srv.user create", err.Error(), 500)
	}
	rsp.Id = 1
	rsp.CreateRequest = req
	return nil
}

// Edit ...
func (h *User) Edit(ctx context.Context, req *user.EditRequest, rsp *user.EditResponse) error {

	return nil
}
