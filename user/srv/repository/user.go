package repository

import (
	"github.com/xionglongjun/micro-mall/user/srv/models"

	"github.com/jinzhu/gorm"
)

// UserRepo ...
type UserRepo struct {
	*gorm.DB
}

// Create ...
func (repo *UserRepo) Create(userModel *models.User) error {
	tx := repo.Begin()

	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
	}()

	if err := tx.Error; err != nil {
    return err
	}
	
	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}