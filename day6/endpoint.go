package main

import (
	"encoding/json"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set header agar respons dalam format JSON
	w.Header().Set("Content-Type", "application/json")

	// Buat response dalam bentuk JSON
	response := map[string]string{"message": "Hello, World!"}

	// Encode response menjadi JSON dan kirim ke client
	json.NewEncoder(w).Encode(response)

}
