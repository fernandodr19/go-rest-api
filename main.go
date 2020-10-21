package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All articles endpoint hit")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Println("Server up");
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Printf("haha\n")
	
	Articles = []Article{
        Article{"Hello", "Article Description", "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	// fmt.Printf("%+v\n", Articles)
	for i, a := range Articles {
		fmt.Printf("%d %+v\n", i, a)
	}

	handleRequest()
}