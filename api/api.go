package api

import (
	"log"
	"net/http"
)

func MapRoutes() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("direct", w, r) })
	http.HandleFunc("/profile-pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("profile", w, r) })

	log.Println("map routes OK")
}
