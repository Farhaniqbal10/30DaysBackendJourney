package main 

import "fmt"

func main {
  //gunakan interface jika tipe data nya berbeda
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

  fmt.Println("data pengguna")
  for i user := range users{
    fmt.Printf("%d. nama: %s, umur: %d, email:%s\n", i+1, user["nama"], user["umur"], user["email"])  
  }
  
  }
