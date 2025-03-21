package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/daniil-oliynyk/go-url-shortener/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}
type config struct {
	addr  string
	cache cacheConfig
}

type cacheConfig struct {
	addr     string
	password string
}

func (app *application) mount() *chi.Mux {
	fmt.Printf("\nHello Go URL Shortener !🚀")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Post("/create-short-url", app.createShortUrl)
	})
	r.Get("/{shortUrl}", app.handleShortUrlRedirect)

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("\nListening at %s", app.config.addr)
	return srv.ListenAndServe()
}
