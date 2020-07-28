package api

import (
	"net/http"
)

func Route() {
	http.HandleFunc("/", Hello)
}
