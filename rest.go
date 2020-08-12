package main
import (
	"fmt"
	"encoding/json"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID	string `json:"id"`
	Title	string `json:"title"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	var books []Book
	books = append(books, Book{ID:"404", Title:"Test title"})
	books = append(books, Book{ID:"405", Title:"Test title"})


	fmt.Println("Get books")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	var books []Book
	books = append(books, Book{ID:"404", Title:"Test title"})
	books = append(books, Book{ID:"405", Title:"Test title"})


	fmt.Println("Get book")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {
	fmt.Println("Server up")
	
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")

	panic(http.ListenAndServe(":8080", r))

    panic("Server down")
}

// curl -i -X GET http://localhost:8080/api/books