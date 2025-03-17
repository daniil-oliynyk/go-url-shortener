package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type UrlCache struct {
	rd *redis.Client
}

func (u *UrlCache) Save(shortUrl string, longURl string, userId string) error {
	err := u.rd.Set(context.Background(), shortUrl, longURl, 6*time.Hour).Err()
	if err != nil {
		fmt.Println("Failed saving key url | Error: ", err, " - shortUrl: ", shortUrl, " - longUrl: ", longURl, "\n")
		return err
	}

	return nil
}

func (u *UrlCache) Retrieve(shortUrl string) (string, error) {

	res, err := u.rd.Get(context.Background(), shortUrl).Result()
	if err != nil {
		fmt.Sprintf("Failed RetrieveInitialUrl url | Error: ", err, "- shortUrl: ", shortUrl, "\n")
	}
	return res, err
}
