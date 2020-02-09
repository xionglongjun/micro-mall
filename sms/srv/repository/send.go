package repository

import (
	"github.com/xionglongjun/micro-mall/sms/srv/models"

	"github.com/jinzhu/gorm"
)

// SendRepo ...
type SendRepo struct {
	*gorm.DB
}

// Create ...
func (repo *SendRepo) Create(sendModel *models.Send) error {
	tx := repo.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&sendModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
