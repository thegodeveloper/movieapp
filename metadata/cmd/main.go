package main

import (
	"github.com/thegodeveloper/movieapp/metadata/internal/controller/metadata"
	httphandler "github.com/thegodeveloper/movieapp/metadata/internal/handler/http"
	"github.com/thegodeveloper/movieapp/metadata/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
