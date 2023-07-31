package handlers

import (
	"URLShortener/pkg/getlink"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// git check
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	log.Println("Get Home")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortLink := r.FormValue("shortLink")
	if shortLink == "" {
		http.Error(w, "Shortened URL is required", http.StatusBadRequest)
		return
	}

	// Implement logic to retrieve the full link from the shortened version
	fullLink, err := getlink.GetFullLink(shortLink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve full link: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Full link: %s", fullLink)
}
