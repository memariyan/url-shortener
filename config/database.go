package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbName   = "url_shortener"
)

func DatabaseConnection() *gorm.DB {

	dsn := "root:root@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"

	/*	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)*/

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
