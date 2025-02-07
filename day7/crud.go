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
