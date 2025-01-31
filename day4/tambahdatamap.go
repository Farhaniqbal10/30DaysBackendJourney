package main 

import "fmt"

func main {
  //gunakan interface jika tipe data nya beda beda
 var users []map[string]interface{}
  for {
    var nama string
    var umur int
    var email string

    fmt.Println("masukan nama")
    fmt.Scanln(&nama)

    fmt.Println("masukan umur")
    fmt.Scanln(&umur)

    fmt.Println("masukan email")
    fmt.Scanln(&email)

    user := map[string]interface{
      "nama": nama,
      "umur": umur,
      "email" : email,
      }

    users = append(users, user)

    var lagi string
    fmt.Println("tambah data lagi ? (y/n) : ")
    fmt.Scanln(&lagi)

    if lagi != "y" {
      break  
      }
    }
  
  }
