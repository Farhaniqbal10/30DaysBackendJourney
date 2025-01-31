package main

import "fmt"

type User struct{
  Id int
  Nama string
  Umur int
  Email string  
  }

func main() {
  var user[] User 

  for {
    var Id int 
    var nama string
    var umur int
    var email string

  
        fmt.Print("Masukkan ID: ")
        fmt.Scanln(&id)

        fmt.Print("Masukkan Nama: ")
        fmt.Scanln(&nama)

        fmt.Print("Masukkan Umur: ")
        fmt.Scanln(&umur)

        fmt.Print("Masukkan Email: ")
        fmt.Scanln(&email)

    users = append(users, User{Id: id, Nama: nama, Umur: umur, Email: email})
  
        var lagi string
        fmt.Print("Tambah lagi? (y/n): ")
        fmt.Scanln(&lagi)

        if lagi != "y" {
            break
        }
    
  }

 fmt.Println("\nData Pengguna:")
    for _, user := range users {
        fmt.Printf("ID: %d | Nama: %s, Umur: %d, Email: %s\n", user.Id, user.Nama, user.Umur, user.Email)
    }

    // Pencarian pengguna berdasarkan ID
    var searchID int
    fmt.Print("\nMasukkan ID pengguna yang ingin dicari: ")
    fmt.Scanln(&searchID)

    var found bool
    for _, user := range users {
        if user.Id == searchID {
            fmt.Println("\nData Pengguna Ditemukan:")
            fmt.Printf("ID: %d | Nama: %s, Umur: %d, Email: %s\n", user.Id, user.Nama, user.Umur, user.Email)
            found = true
            break
      
        }
    }

    if !found {
        fmt.Println("ID tidak ditemukan!")
    }

    var carilagi string 
          fmt.Println("cari lagi ? (y/n) :")
          fmt.Scanln(&carilagi)
          if lagi != "y" {
            break
            }


}
