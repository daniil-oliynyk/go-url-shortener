package store

import (
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Urls interface {
		Save(shortUrl string, longUrl string, userId string) error
		Retrieve(shortUtl string) (string, error)
	}
}

func NewRedisStore(rd *redis.Client) Storage {
	return Storage{
		Urls: &UrlCache{rd},
	}
}
