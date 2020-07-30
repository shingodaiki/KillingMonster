package api

import (
	"log"
	"net/http"
)

func MapRoutes() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("direct", w, r) })
	http.HandleFunc("/github-pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("github", w, r) })

	log.Println("map routes OK")
}
