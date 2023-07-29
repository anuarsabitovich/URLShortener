package main

import (
	"URLShortener/web"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started at http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", web.Router()))
}
