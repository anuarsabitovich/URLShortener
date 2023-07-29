package web

import (
	"URLShortener/internal/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/getLink", handlers.GetLink)
	return mux
}
