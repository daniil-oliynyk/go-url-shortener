package cache

import (
	"github.com/redis/go-redis/v9"
)

func New(addr string, password string) *redis.Client {

	rd := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return rd
}
