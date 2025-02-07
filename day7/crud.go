package main

import (
"fmt"
"net/http"
"github.com/gorilla/mux"
)

type Book struct {
  ID  string `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
  }

var books = make(map[string]Book)

func getBooks(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  var bookList []Book
  for _, book := range books {
      bookList = append(bookList, book)
    }
  json.NewEncoder(w).Encode(bookList)
  }

func createBook(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  var book Book
  if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
    }
  books[book.ID] = book
  json.NewEncoder(w).Encode(book)
  }

func deleteBook(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  if _, exists := books[id]; !exists {
    http.Error(w, "Book not found", http.StatusNotFound)
    return
    }
  delete(books, id)
  w.WriterHeader(http.StatusNoContent)
  }

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/books", getBooks).Methods("GET")
  }

  }
