package main

import (
	"github.com/darwinz/kudo-oos/pkg/storage"
	web "github.com/darwinz/kudos/pkg/http"
	"log"
	"net/http"
	"os"
)

func main() {
	httpPort := os.Getenv("PORT")

	repo := storage.NewMongoRepository()
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}
