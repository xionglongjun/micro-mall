package models

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model
	Name     string
	Password string
	Salt     string `gorm:"default:'******';not null"`
	Mobile   string `gorm:"type:varchar(50);unique_index;not null"`
	Enabled  uint   `gorm:"type:tinyint(1);default:1;not null"`
}
