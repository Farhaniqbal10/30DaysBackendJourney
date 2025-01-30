package main 

import "fmt"

func main(){
  mahasiswa := make(map[string]string)

  mahasiswa["1001"] = "Farhan"
  mahasiswa["1002"] = "Maulana"
  mahasiswa["1003"] = "Iqbal"

  fmt.Println(mahasiswa)

  fmt.Println(mahasiswa["1001"])

  for nim, nama := range mahasiswa {
  fmt.Println(a...: "Nim", nim, "bernama", nama)  
  }
}
