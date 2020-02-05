package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Code ...
type Code struct {
	gorm.Model
	SendID uint
	Send Send
	UserID uint `gorm:"index"`
	Mobile   string `gorm:"index;type:varchar(50);not null"`
	Code string `gorm:"type:varchar(50);not null"`
	Use  uint   `gorm:"type:tinyint(1);default:1;not null"`
	ExpiresAt time.Time `gorm:"type:datetime;not null"`
}
