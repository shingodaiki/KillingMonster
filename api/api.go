package api

import (
	"log"
	"net/http"
)

func MapRoutes() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pixel.gif", pixel)

	log.Println("map routes OK")
}
