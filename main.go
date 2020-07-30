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
		log.Fatal("$PORT must be set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	api.InitStore(dbURL)
	defer api.CloseStore()

	mKey := os.Getenv("MASTER_KEY")
	if mKey == "" {
		log.Fatal("$MASTER_KEY must be set")
	}
	api.InitStats(mKey)

	api.MapRoutes()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
