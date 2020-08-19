package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaycub/go-rest/connect"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome Home")
	fmt.Println("endpoint hit: homepage")
}

func createArticle(response http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)

	json.NewEncoder(response).Encode(article)
}

func returnAllArticles(response http.ResponseWriter, request *http.Request) {
	fmt.Println("endpoint hit: all articles")
	json.NewEncoder(response).Encode(articles)
}

func returnSingleArticle(response http.ResponseWriter, request *http.Request) {
	key := mux.Vars(request)["id"]
	fmt.Println("endpoint hit: article by id")
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(response).Encode(article)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//GET
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	//POST
	myRouter.HandleFunc("/article", createArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":7890", myRouter))
}

func main() {
	articles = []Article{
		Article{Id: "1", Title: "Title 1", Desc: "Description 1", Content: "Content 1"},
		Article{Id: "2", Title: "Title 2", Desc: "Description 2", Content: "Content 2"},
	}

	fmt.Println("listening on port 7890")
	connect.Connect()
	handleRequests()
}
