package main 

import ( 
"fmt"
"net/http" 
)

func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo ini adalah server http dengan go")
}

func main(){

	http.HandleFunc("/", handler)

	fmt.Println("server berjalan di localhost")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("error saat menjalankan server", err)
	}
}

