package database

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"url-shortner/internal/config"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	return redisClient
}

func ConnectRedis(config *config.Redis) *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	redisClient = client

	return redisClient
}
