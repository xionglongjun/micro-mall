package models

import "github.com/jinzhu/gorm"

// Send ...
type Send struct {
	gorm.Model
	UserID uint `gorm:"index"`
	Provider string `gorm:"type:varchar(100);not null"`
	Mobile   string `gorm:"type:varchar(50);not null"`
	Content string `gorm:"type:text;not null"`
	Success  uint   `gorm:"type:tinyint(1);default:0;not null"`
}
