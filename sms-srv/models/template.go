package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Template ...
type Template struct {
	gorm.Model
	Mode      uint      `gorm:"type:tinyint(1);default:0;not null"`
	Provider  string    `gorm:"type:varchar(100);not null"`
	Sign      string    `gorm:"type:varchar(100);not null"`
	Content   string    `gorm:"type:text;not null"`
	BizType   uint      `gorm:"type:tinyint(1);default:0;index;not null"`
	ExpiresAt time.Time `gorm:"type:datetime"`
	Enabled   uint      `gorm:"type:tinyint(1);default:1;index;not null"`
}

// ScopeMode scopes
func (t *Template) ScopeMode(mode TempMode) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("mode = ?", uint(mode))
	}
}

// TempBizType ...
type TempBizType int

const (
	// TempLogin ...
	TempLogin TempBizType = iota
	// TempRegister ...
	TempRegister
	// TempEditPassword ...
	TempEditPassword
	// TempResetPassword ...
	TempResetPassword
)

func (e TempBizType) String() string {
	str := ""
	switch e {
	case TempRegister:
		str = "register"
	case TempEditPassword:
		str = "editPassword"
	case TempResetPassword:
		str = "resetPassword"
	default:
		str = "login"
	}
	return str
}

// TempMode ...
type TempMode int

const (
	// TempVerification ...
	TempVerification TempMode = iota
	// TempNotice ...
	TempNotice
)

func (e TempMode) String() string {
	str := ""
	switch e {
	case TempNotice:
		str = "notice"
	default:
		str = "verification"
	}
	return str
}
