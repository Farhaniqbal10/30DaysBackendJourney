package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/gorilla/mux
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set header agar respons dalam format JSON
	w.Header().Set("Content-Type", "application/json")

	// Buat response dalam bentuk JSON
	response := map[string]string{"message": "Hello, World!"}

	// Encode response menjadi JSON dan kirim ke client
	json.NewEncoder(w).Encode(response)

}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	vars := mux.Vars(r)
	name := vars["name"]

	response := map[string]string{"message" : "hello" + name + "!"}
	json.NewEncoder(w).Encode(response)
	
}


func main() {
	r.HandleFunc("/greet/", greetHandler).Methods("GET")
	http.HandleFunc("/hello", helloHandler) // Endpoint /hello
	http.ListenAndServe(":3000", nil)       // Jalankan server di port 3000
}
