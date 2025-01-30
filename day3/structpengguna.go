package main 

import "fmt"

type Pengguna struct {
  ID,Nama,Email string
  }

func main () {
 var pengguna Pengguna
  pengguna.ID = "1001"
  pengguna.Nama = "geboy"
  pengguna.Email = "geboy123@gmail.com"

  fmt.Println(pengguna)

  //penulisan lain untuk struct
  pengguna := Pengguna {
    ID = "1001",
    Nama = "farhan",
    email = "farhaniqbal10@gmail.com",
    }
  }

