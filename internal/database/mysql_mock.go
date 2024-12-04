package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockDB() sqlmock.Sqlmock {
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: mockDB, SkipInitializeWithVersion: true}))
	if err != nil {
		log.Fatal(err)
	}
	mySQL = db
	return sqlMock
}
