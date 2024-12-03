package database

import (
	"github.com/go-redis/redismock/v9"
	"url-shortner/internal/database"
)

func MockRedis() redismock.ClientMock {
	db, mock := redismock.NewClientMock()
	database.Redis = db
	return mock
}
