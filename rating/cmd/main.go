package main

import (
	"github.com/thegodeveloper/movieapp/rating/internal/controller/rating"
	httphandler "github.com/thegodeveloper/movieapp/rating/internal/handler/http"
	"github.com/thegodeveloper/movieapp/rating/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
