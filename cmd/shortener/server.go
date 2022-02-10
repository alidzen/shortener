package main

import (
	"github.com/alidzen/shortener/internal/app"
	"io"
	"log"
	"net/http"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	log.Println(id)
	if id == "" {
		Shortener(w, r)
		return
	}
	LongHandler(w, r)
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Allow only POST", http.StatusBadRequest)
		return
	}
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set up header
	w.Header().Set("content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	shortLink := app.GenerateShortLink(string(b))

	urls[shortLink] = string(b)
	log.Println(urls)

	w.Write([]byte(shortLink))
}

func LongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Allow only GET", http.StatusBadRequest)
		return
	}

	id := getId(r)
	if _, ok := urls[id]; !ok {
		http.Error(w, "The url with provided id was not found", http.StatusNotFound)
		return
	}

	w.Write([]byte(urls[id]))
}
