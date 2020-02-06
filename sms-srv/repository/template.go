package repository

import (
	"sms-srv/enums"
	"sms-srv/models"

	send "sms-srv/proto/send"

	"github.com/jinzhu/gorm"
)

// TemplateRepo ...
type TemplateRepo struct {
	*gorm.DB
}

// Create ...
func (r *TemplateRepo) Create(templateModel *models.Template) error {
	tx := r.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&templateModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetCodeTemplateFirst ...
func (r *TemplateRepo) GetCodeTemplateFirst(provider string, bizType uint) (*models.Template, error) {
	var (
		templateModel models.Template
		err           error
	)
	templateModel.Provider = provider
	r.DB = r.Scopes(models.Scope("enabled", uint(enums.Yes)), models.Scope("mode", uint(send.Mode_Code)), models.Scope("biz_type", bizType))
	if err = r.Where(&templateModel).First(&templateModel).Error; err != nil {
		return nil, err
	}
	return &templateModel, nil
}
