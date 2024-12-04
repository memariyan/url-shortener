package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"url-shortner/internal/config"
)

var mySQL *gorm.DB

func GetDB() *gorm.DB {
	return mySQL
}

func ConnectDB(config *config.MySQL) *gorm.DB {
	if mySQL != nil {
		return mySQL
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mySQL = db

	return db
}
