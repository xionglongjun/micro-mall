package repository

import (
	"sms-srv/models"

	"github.com/jinzhu/gorm"
)

// CodeRepo ...
type CodeRepo struct {
	*gorm.DB
}

// Name ...
func (r *CodeRepo) Name(name string) (*models.Code, error) {
	var (
		codeModel models.Code
	)
	if err := r.Where(codeModel).First(&codeModel).Error; err != nil {
		return nil, err
	}
	return &codeModel, nil
}

// Mobile ...
func (r *CodeRepo) Mobile(mobile string) (*models.Code, error) {
	var (
		codeModel models.Code
	)
	if err := r.Where(codeModel).First(&codeModel).Error; err != nil {
		return nil, err
	}
	return &codeModel, nil
}
