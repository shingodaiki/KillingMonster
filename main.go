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

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("$DATABASE_URL was not set, PgDB will not be available")
	} else {
		api.PgDB(dbURL)
	}

	api.Route()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
