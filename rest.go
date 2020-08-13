package main
import (
	"fmt"
	"encoding/json"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

type GetResponse struct {
	Succes	  bool `json:"success"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	Book *Book `json:"book,omitempty"`
	Books *[]Book `json:"books,omitempty"`
}

type Book struct {
	ID	string `json:"id"`
	Title	string `json:"title"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("Get books")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetResponse{true, "","", nil, &books})

	/*
	curl -i -X GET http://localhost:8080/api/books
	*/
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get book")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	paramId := params["id"]

	for _, book := range books {
		if book.ID == paramId {
			json.NewEncoder(w).Encode(GetResponse{true, "","", &book, nil})
			return
		}
	}
	json.NewEncoder(w).Encode(GetResponse{false, "404","Not found", nil, nil})

	/*
	curl -i -X GET http://localhost:8080/api/book/404
	*/
}

func createBook(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("Create book")

	w.Header().Set("Content-Type", "application/json")
	var newBook Book

	_ = json.NewDecoder(r.Body).Decode(&newBook)
	fmt.Println(newBook)
	newBook.ID = "10" //Don't do it at home folks
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook) // Create custom response here?

	/*
	curl --location --request POST 'http://localhost:8080/api/books/create' \
	--header 'Content-Type: application/json' \
	--data-raw '{
		"title": "Test title"
	}'
	*/
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update book")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	paramId := params["id"]

	for index, book := range books {
		if book.ID == paramId {
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = book.ID
			books[index] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return 
		}
	}
	json.NewEncoder(w).Encode(GetResponse{false, "404","Not found", nil, nil})

	/*
	curl --location --request PUT 'http://localhost:8080/api/books/update/405' \
	--header 'Content-Type: application/json' \
	--data-raw '{
		"title": "New test title"
	}'
	*/
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete book")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	paramId := params["id"]

	for index, book := range books {
		if book.ID == paramId {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(GetResponse{true, "","", nil, nil})
			return
		}
	}
	json.NewEncoder(w).Encode(GetResponse{false, "404","Not found", nil, nil})

	/*
	curl -i -X DELETE http://localhost:8080/api/books/delete/404
	*/
}

func main() {
	fmt.Println("Server up")
	
	books = append(books, Book{ID:"404", Title:"Test title"})
	books = append(books, Book{ID:"405", Title:"Test title"})

	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/create", createBook).Methods("POST")
	r.HandleFunc("/api/books/update/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/delete/{id}", deleteBook).Methods("DELETE")

	panic(http.ListenAndServe(":8080", r))

    panic("Server down")
}

// curl -i -X GET http://localhost:8080/api/books