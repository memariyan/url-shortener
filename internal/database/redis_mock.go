package database

import (
	"github.com/go-redis/redismock/v9"
)

func MockRedis() redismock.ClientMock {
	db, mock := redismock.NewClientMock()
	redisClient = db
	return mock
}
