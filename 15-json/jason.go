package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Auther string `json:"auther"`
	Amount int    `json:"amount"`
}

var books = []Book{
	{"1", "Deep Work", "Cal Newport", 34},
	{"2", "Atomic Habits", "James Clear", 40},
	{"3", "The Psychology of Money", "Morgan Housel", 28},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", getBooks)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
