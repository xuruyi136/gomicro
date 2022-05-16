package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	//默认结构体
	gorm.Model
	Username string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(100)"`
	Channel  int64  `gorm:"type:tinyint(2)"`
}
