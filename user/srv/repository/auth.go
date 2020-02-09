package repository

import (
	"github.com/xionglongjun/micro-mall/user/srv/models"

	"github.com/jinzhu/gorm"
)

// AuthRepo ...
type AuthRepo struct {
	*gorm.DB
}

// Name ...
func (r *AuthRepo) Name(name string) (*models.User, error) {
	var (
		userModel models.User
	)
	userModel.Name = name
	userModel.Enabled = 1
	if err := r.Where(userModel).First(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}

// Mobile ...
func (r *AuthRepo) Mobile(mobile string) (*models.User, error) {
	var (
		userModel models.User
	)
	userModel.Mobile = mobile
	userModel.Enabled = 1
	if err := r.Where(userModel).First(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}
