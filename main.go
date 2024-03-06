package main

import (
	"log"
	"net/http"
	"url-shortener/templates"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", templ.Handler(templates.Home("")))

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		getShorten(w, r)
	})

	log.Println("Starting on port 3333")
	http.ListenAndServe(":3333", nil)
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("url")
	templates.Home(longUrl).Render(r.Context(), w)
}
