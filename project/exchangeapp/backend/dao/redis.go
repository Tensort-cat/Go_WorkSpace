package dao

import (
	"CurrencyExchangeApp/config"
	"log"

	"github.com/go-redis/redis"
)

var (
	RedisTemplate *redis.Client
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.Host + config.AppConfig.Redis.Port,
		DB:       config.AppConfig.Redis.DB,
		Password: config.AppConfig.Redis.Password,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to redis")
	}
	RedisTemplate = RedisClient
}
