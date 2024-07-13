package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:10As2000*@tcp(127.0.0.1)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
