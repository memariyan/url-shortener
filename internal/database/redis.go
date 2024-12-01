package database

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"url-shortner/internal/config"
)

var Redis *redis.Client

func ConnectRedis(config *config.Redis) *redis.Client {
	if Redis != nil {
		return Redis
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
	Redis = client

	return Redis
}
