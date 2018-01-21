package main

import (
	"html/template"
	"net/http"
	"strconv"
	"log"
)

var templates = template.Must(template.ParseFiles(
	"index.html",
))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	afterStr := r.URL.Query().Get("after")
	after, err := strconv.Atoi(afterStr)
	if err != nil {
		log.Println("Could not parse after date", err)
		after = 0
	}
	posts, err := getPosts(after)
	if err != nil {
		log.Println("Could not get posts:", err)
		return
	}
	templates.ExecuteTemplate(w, "index.html", posts)
}

func main() {
	// Register handlers
	http.HandleFunc("/", indexHandler)

	// Serve forever
	log.Fatal(http.ListenAndServe(":8080", nil))
}
