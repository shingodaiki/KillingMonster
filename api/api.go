package api

import (
	"net/http"
)

func PgDB(url string) {
}

func Route() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pixel.gif", pixel)
}
