package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	fmt.Println("endpoint hit: article by id")
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":7890", myRouter))
}

func main() {
	articles = []Article{
		Article{Id: "1", Title: "Title 1", Desc: "Description 1", Content: "Content 1"},
		Article{Id: "2", Title: "Title 2", Desc: "Description 2", Content: "Content 2"},
	}

	fmt.Println("listening on port 7890")
	handleRequests()
}
