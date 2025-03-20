package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/daniil-oliynyk/go-url-shortener/internal/shortener"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

type UrlCreationPayload struct {
	LongUrl string `schema:"long_url,required"`
	UserId  string `schema:"user_id,required"`
}

func (app *application) createShortUrl(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.String())
	var payload UrlCreationPayload
	var decoder = schema.NewDecoder()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&payload, u.Query())
	if err != nil {
		// Handle error
		fmt.Println("err:", err)
	}
	// fmt.Println(u.Query())
	// fmt.Println("longurl", payload.LongUrl, "userid", payload.UserId)
	shortUrl := shortener.GenerateShortLink(payload.LongUrl, payload.UserId)

	err = app.store.Urls.Save(shortUrl, payload.LongUrl, payload.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("short url not created"))
		return
	}
	host := "http://www.localhost:8080/"
	type SuccessResponse struct {
		Msg      string
		ShortUrl string
	}
	resp := SuccessResponse{"short url created", host + shortUrl}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (app *application) handleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	fmt.Println(shortUrl)
	longUrl, err := app.store.Urls.Retrieve(shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("url not found"))
		return
	}

	type SuccessResponse struct {
		LongUrl string
	}
	resp := SuccessResponse{longUrl}
	w.WriteHeader(http.StatusPermanentRedirect)
	json.NewEncoder(w).Encode(resp)

	http.Redirect(w, r, longUrl, http.StatusFound)

}
