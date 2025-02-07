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
