package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
	fmt.Println("endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: all articles")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":7890", nil))
}

func main() {
	articles = []Article{
		Article{Title: "Title 1", Desc: "Description 1", Content: "Content 1"},
		Article{Title: "Title 2", Desc: "Description 2", Content: "Content 2"},
	}
	fmt.Println("listening on port 7890")
	handleRequests()
}
