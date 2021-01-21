package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Desc   string `json:"Desc"`
	UID    string `json:"UID"`
}

var articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Article management system: homePage()")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getArticles()")

	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteArticleAtUid(params["uid"])

	json.NewEncoder(w).Encode(articles)
}

func _deleteArticleAtUid(uid string) {
	for index, article := range articles {
		if article.UID == uid {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)

	params := mux.Vars(r)

	_deleteArticleAtUid(params["uid"])

	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", postArticle).Methods("POST")
	router.HandleFunc("/articles/{uid}", deleteArticle).Methods("POST")
	router.HandleFunc("/articles/{uid}", deleteArticle).Methods("POST")
	router.HandleFunc("/articles/{uid}", updateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	articles = append(articles, Article{
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "Go Programming tutorial",
		UID:    "001",
	})

	articles = append(articles, Article{
		Title:  "Python Tricks",
		Author: "Pythonista",
		Desc:   "Python Programming tutorial",
		UID:    "002",
	})

	handleRequests()
}
