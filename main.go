package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/shorten", getShorten)
	log.Println("Starting on port 3333")

	http.ListenAndServe(":3333", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
		<h1>URL Shortener</h1>
		<form action="/shorten" method="post">
			<input placeholder="Your long URL" name="url">
			<button>Shorten</button>
		</form>
	`)
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("url")
	io.WriteString(w, longUrl)
}
