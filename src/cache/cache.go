package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var Ctx = context.TODO()

func newDatabaseClient() (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func GetString(key string) string {
	db, err := newDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	return db.Client.Get(Ctx, key).Val()
}

func SetString(key string, value string) {
	db, err := newDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	db.Client.Set(Ctx, key, value, 0)
}
