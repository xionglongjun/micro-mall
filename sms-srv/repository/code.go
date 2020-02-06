package repository

import (
	"sms-srv/enums"
	"sms-srv/models"
	"time"

	"github.com/jinzhu/gorm"
)

// CodeRepo ...
type CodeRepo struct {
	*gorm.DB
	model models.Code
}

// Create ...
func (repo *CodeRepo) Create(sendModel *models.Code) error {
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

// Validate ...
func (r *CodeRepo) Validate(mobile, code string, bizType uint) (*models.Code, error) {
	var (
		codeModel models.Code
		now       = time.Now()
	)

	db := r.Scopes(models.Scope("is_use", uint(enums.No))).
		Where("biz_type = ?", bizType).Where("expires_at > ?", now)
	codeModel.Mobile = mobile
	codeModel.Code = code
	if err := db.Where(codeModel).First(&codeModel).Error; err != nil {
		return nil, err
	}
	if err := r.UseCode(codeModel.ID); err != nil {
		return nil, err
	} else {
		codeModel.IsUse = uint(enums.Yes)
	}

	return &codeModel, nil
}

func (r *CodeRepo) UseCode(id uint) error {
	var codeModel models.Code
	codeModel.ID = id
	return r.Model(codeModel).Update("is_use", uint(enums.Yes)).Error
}

// Mobile ...
func (r *CodeRepo) Mobile(mobile string) (*models.Code, error) {
	var (
		codeModel models.Code
	)
	codeModel.Mobile = mobile
	if err := r.Where(codeModel).First(&codeModel).Error; err != nil {
		return nil, err
	}
	return &codeModel, nil
}
