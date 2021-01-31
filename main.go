package main

//Basic imports
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"Isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

//Intialize books slice
var books []Book

//Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get book with id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for _,item := range books {
		if(item.ID == params["id"]){
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create book with post method
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update books with id with PUT method
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var book Book
			_= json.NewDecoder(r.Body).Decode(&book)
	for _,item := range books {
		if(item.ID == params["id"]) {
			book.ID = strconv.Itoa(rand.Intn(10000000))
			item = book
			json.NewEncoder(w).Encode(item)
			return
		}
		break
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Delete Books with ID using DELETE methos
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index] , books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Intializing router
	r := mux.NewRouter()

	//Intsializing mock data for books
	books = append(books, Book{ID : "1" , Isbn : "1265415" , Title: "Reach your goals" , Author : &Author{firstname: "john" , lastname : "doe"}})
	books = append(books, Book{ID : "2" , Isbn : "8646545" , Title: "Read my Mind" , Author : &Author{firstname: "Mike" , lastname : "Vosaski"}})
	books = append(books, Book{ID : "3" , Isbn : "1485238" , Title: "Black Balls" , Author : &Author{firstname: "Kim" , lastname : "Kardashian"}})
	//Adding Routes
	r.HandleFunc("/api/getbooks" , getBooks).Methods("GET")
	r.HandleFunc("/api/getbook/{id}" , getBook).Methods("GET")
	r.HandleFunc("/api/createbook" , createBook).Methods("POST")
	r.HandleFunc("/api/updatebook/{id}" , updateBook).Methods("PUT")
	r.HandleFunc("/api/deletebook/{id}" , deleteBook).Methods("DELETE")

	//Run server
	log.Fatal(http.ListenAndServe(":3000" , r))
}