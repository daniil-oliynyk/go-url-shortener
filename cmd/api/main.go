package main

import (
	"context"
	"fmt"
	"log"

	"github.com/daniil-oliynyk/go-url-shortener/internal/cache"
	"github.com/daniil-oliynyk/go-url-shortener/internal/store"
)

func main() {

	cfg := config{
		addr: ":8080",
		cache: cacheConfig{
			addr:     ":6379",
			password: "",
		},
	}

	cache := cache.New(cfg.cache.addr, cfg.cache.password)

	pong, err := cache.Ping(context.Background()).Result()
	if err != nil {
		fmt.Sprintf("Error init Redis: %v", err)
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

	redisStore := store.NewRedisStore(cache)

	app := &application{
		config: cfg,
		store:  redisStore,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
