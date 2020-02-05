package repository

import (
	"fmt"
	"sms-srv/enums"
	"sms-srv/models"

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
func (r *TemplateRepo) GetCodeTemplateFirst(provider string, bizType models.TempBizType) (*models.Template, error) {
	var (
		templateModel models.Template
	)

	templateModel.Mode = models.Verification
	templateModel.Enabled = enums.Yes
	templateModel.Provider = provider
	templateModel.BizType = bizType
	fmt.Println(enums.Yes)
	if err := r.Where(templateModel).First(&templateModel).Error; err != nil {
		return nil, err
	}
	return &templateModel, nil
}
