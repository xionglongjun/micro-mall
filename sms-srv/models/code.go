package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Code ...
type Code struct {
	gorm.Model
	SendID    uint
	UserID    uint      `gorm:"index"`
	Mobile    string    `gorm:"index;type:varchar(50);not null"`
	Code      string    `gorm:"type:varchar(50);not null"`
	BizType   uint      `gorm:"type:tinyint(1);default:0;index;not null"`
	IsUse     uint      `gorm:"type:tinyint(1);default:0;not null"`
	ExpiresAt time.Time `gorm:"type:datetime;not null"`
	Send      Send
}
