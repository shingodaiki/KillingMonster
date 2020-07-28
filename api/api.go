package api

import (
	"net/http"
)

func Route() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pixel.gif", pixel)
}
