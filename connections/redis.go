package connections

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisConn *RedisConnection
type RedisConnection struct {
	Client *redis.Client
}

type redisConfig struct {
	Address  string
	Password string
	DB       int
}



func newRedisConnection(cfg redisConfig) (*RedisConnection, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis")

	return &RedisConnection{Client: rdb}, nil
}


func (rc *RedisConnection) Close() {
	err := rc.Client.Close()
	if err != nil {
		log.Println("Error disconnecting from Redis:", err)
	}
}


func LoadRedisConfig() redisConfig {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal("Failed to parse REDIS_DB environment variable")
	}

	redisConfig := redisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
	return redisConfig
}

func InitRedis() {
	var err error
	RedisConn, err = newRedisConnection(LoadRedisConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
