package main

import "fmt"

type user struct{
  Id int
  nama string
  umur int
  email string  
  }

func main() {
  var user[] User 

  for {
    var Id int 
    nama string
    umur int
    }

  
        fmt.Print("Masukkan ID: ")
        fmt.Scanln(&id)

        fmt.Print("Masukkan Nama: ")
        fmt.Scanln(&nama)

        fmt.Print("Masukkan Umur: ")
        fmt.Scanln(&umur)

        fmt.Print("Masukkan Email: ")
        fmt.Scanln(&email)

    user = append(users, User{ID: id, Nama: nama, Umur: umur, Email: email})
  
        var lagi string
        fmt.Print("Tambah lagi? (y/n): ")
        fmt.Scanln(&lagi)

        if lagi != "y" {
            break
        }
    
  }

 fmt.Println("\nData Pengguna:")
    for _, user := range users {
        fmt.Printf("ID: %d | Nama: %s, Umur: %d, Email: %s\n", user.ID, user.Nama, user.Umur, user.Email)
    }
