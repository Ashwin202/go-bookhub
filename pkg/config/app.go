package config

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	os.Getenv("DATABASE_USER"),
	os.Getenv("DATABASE_PASSWORD"),
	os.Getenv("DATABASE_URL"),
	os.Getenv("DATABASE_NAME"))
	d, err := gorm.Open("mysql", databaseURI)
	if err != nil {
		panic(err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
