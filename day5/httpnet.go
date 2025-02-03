package main 

import ( 
"fmt"
"net/http" 
)

func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "halo ini adalah server http dengan go")
}

func main(){


}

