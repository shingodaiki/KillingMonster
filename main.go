package main

import (
	"log"
	"net/http"
	"os"

	"mfatihmar/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT was not set, using default port 8080")
		port = "8080"
	}

	api.Route()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
