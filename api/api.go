package api

import (
	"log"
	"net/http"
)

func MapRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { pixel("direct", w, r) })
	http.HandleFunc("/about-pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("about", w, r) })
	http.HandleFunc("/blog-pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("blog", w, r) })
	http.HandleFunc("/post-pixel.gif", func(w http.ResponseWriter, r *http.Request) { pixel("post", w, r) })

	log.Println("map Routes OK")
}
