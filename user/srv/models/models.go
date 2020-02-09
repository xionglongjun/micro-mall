package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Driver ...
type Driver interface {
	Connection() (*gorm.DB, error)
}

// Mysql ...
type Mysql struct {
	Host string
	User string
	Password string
	Name string
	Charset string
	Prefix string
	Debug bool
}

// Connection ...
func (m *Mysql) Connection() (*gorm.DB, error) {
	if m.Charset == "" {
		m.Charset = "utf8mb4"
	}

	if m.Host == "" {
		m.Host = "localhost:3306"
	}
	
	if m.Name == "" {
		m.Name = "root"
	}

	if m.Password == "" {
		m.Password = "123456"
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Name, m.Charset))


	if err != nil {
		log.Fatalf("models.Connection err: %v", err)
		return nil, err
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return m.Prefix + defaultTableName;
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.LogMode(m.Debug)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db, err
}