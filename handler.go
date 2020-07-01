package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Error struct {
	ErrorCode   int    `json:"error"`
	Message     string `json:"message"`
	Remediation string `json:"remediation"`
}

type Book struct {
	Id     int    `json:"ID"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Isbn   string `json:"isbn"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book = []Book{}

func main() {
	fmt.Println("Getting Started with the Golang Books Management Rest API.")
	handleHttpRequests()
}

func handleHttpRequests() {
	//	http.HandleFunc("/", handleRootAPI)
	//	log.Fatal(http.ListenAndServe(":8081", nil))

	// Register dummy data into Books

	books = append(books, Book{Id: 1, Title: "5 Point Someone", Author: Author{FirstName: "Chetan", LastName: "Bhagat"}, Isbn: "1001"})
	books = append(books, Book{Id: 2, Title: "I too had a love story", Author: Author{FirstName: "Chetan", LastName: "Bhagat"}, Isbn: "1001"})
	router := mux.NewRouter()
	router.HandleFunc("/", handleRootAPI).Methods("GET")
	router.HandleFunc("/books", handleGEBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handleGETBook).Methods("GET")
	router.HandleFunc("/books", handleCreatebooks).Methods("POST")
	router.HandleFunc("/books/{id}", handleUpdatebooks).Methods("PUT")
	router.HandleFunc("/books/{id}", handleDeletebooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func handleRootAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root API has been invoked.")
}

func handleGEBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Collection GET Book API has been invoked.")
	fmt.Println(books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func handleGETBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single GET Book API has been invoked.")
	params := mux.Vars(r)
	bookId, error := strconv.Atoi(params["id"])
	if error != nil {
		error := Error{ErrorCode: http.StatusBadRequest, Message: "Invalid Id format.", Remediation: "Usage: Please provide numeric Id"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	for _, value := range books {
		if value.Id == bookId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	err := Error{ErrorCode: http.StatusNotFound, Message: "Book not found.", Remediation: "Usage: Please provide valid Book Id"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)
	return
}

func handleCreatebooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Book API has been invoked.")
	var book Book
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = rand.Intn(1000000)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func handleUpdatebooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update API has been invoked.")
	params := mux.Vars(r)
	bookId, error := strconv.Atoi(params["id"])
	if error != nil {
		errorObj := Error{ErrorCode: http.StatusBadRequest, Message: "Invalid Id Format", Remediation: "Usage: Please provide valid numeric Id."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorObj)
	}
	for index, book := range books {
		if bookId == book.Id {
			books = append(books[:index], books[index+1:]...)
			var updatedBook Book
			json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.Id = rand.Intn(1000000)
			books = append(books, updatedBook)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	err := Error{ErrorCode: http.StatusNotFound, Message: "Book not found.", Remediation: "Usage: Please provide valid book Id."}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)

}

func handleDeletebooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete API has been invoked.")
	params := mux.Vars(r)
	bookId, error := strconv.Atoi(params["id"])
	if error != nil {
		errorObj := Error{ErrorCode: http.StatusBadRequest, Message: "Invalid Id Format", Remediation: "Usage: Please provide valid numeric Id."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorObj)
		return
	}
	for index, book := range books {
		if bookId == book.Id {
			books = append(books[:index], books[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	err := Error{ErrorCode: http.StatusNotFound, Message: "Book not found.", Remediation: "Usage: Please provide valid book Id."}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)

}
