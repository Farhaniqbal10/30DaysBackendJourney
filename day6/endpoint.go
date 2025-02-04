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


func main() {
	http.HandleFunc("/hello", helloHandler) // Endpoint /hello
	http.ListenAndServe(":3000", nil)       // Jalankan server di port 3000
}
