package main

import (
	"encoding/json"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set header agar respons dalam format JSON
	w.Header().Set("Content-Type", "application/json")

}
