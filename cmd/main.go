package main

import (
	"fmt"
	web "github.com/darwinz/kudos/pkg/http"
	"github.com/darwinz/kudos/pkg/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	httpPort := os.Getenv("KUDOS_PORT")
	httpHost := os.Getenv("KUDOS_HOST")

	repo := storage.NewMongoRepository()
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", httpHost, httpPort), webService.Router))
}
