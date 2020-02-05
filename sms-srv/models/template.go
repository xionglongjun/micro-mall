package models

import (
	"sms-srv/enums"
	"time"

	"github.com/jinzhu/gorm"
)

// Template ...
type Template struct {
	gorm.Model
	Mode      TempMode      `gorm:"type:tinyint(1);default:0;not null"`
	Provider  string        `gorm:"type:varchar(100);not null"`
	Sign      string        `gorm:"type:varchar(100);not null"`
	Content   string        `gorm:"type:text;not null"`
	BizType   TempBizType   `gorm:"type:tinyint(1);default:0;index;not null"`
	ExpiresAt time.Time     `gorm:"type:datetime"`
	Enabled   enums.Enabled `gorm:"type:tinyint(1);default:1;index;not null"`
}

// TempBizType ...
type TempBizType uint

const (
	// Login ...
	Login TempBizType = iota
	// Register ...
	Register
	// EditPassword ...
	EditPassword
	// ResetPassword ...
	ResetPassword
)

func (e TempBizType) String() string {
	str := ""
	switch e {
	case Register:
		str = "register"
	case EditPassword:
		str = "editPassword"
	case ResetPassword:
		str = "resetPassword"
	default:
		str = "login"
	}
	return str
}

// TempMode ...
type TempMode uint

const (
	// Verification ...
	Verification TempMode = iota
	// Notice ...
	Notice
)

func (e TempMode) String() string {
	str := ""
	switch e {
	case Notice:
		str = "notice"
	default:
		str = "verification"
	}
	return str
}
