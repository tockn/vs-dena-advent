package main

import (
	"log"
	"net/http"

	"github.com/tockn/vs-dena-advent/api/router"
	"github.com/tockn/vs-dena-advent/api/server"
	"github.com/tockn/vs-dena-advent/persistence/memory"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	likeRepo := memory.NewLikesRepository()
	s := server.NewServer(likeRepo)
	r := router.New(s)
	log.Println("serving...")
	return http.ListenAndServe(":8080", r)
}