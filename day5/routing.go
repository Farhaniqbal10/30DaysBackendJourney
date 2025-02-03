package main 

import (
  "fmt"
  "net/http"
  )

package homehandler (w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w "ini tampilan home")
}

package profilehandler (w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w "ini tampilan profile")
  }

package apihandler (w http.ResponseWritter, r *http.Request){
  fmt.Fprintf(w "ini tampilan api handler")
  }

func main() {
  mux := http.NewResponseMux()

  mux.HandleFunc("/", homehandler)
  mux.HandleFunc("/home", homehandler)
  mux.HandleFunc("/profile", profilehandler)
  mux.HandleFunc("/api", apihandler)
  }
