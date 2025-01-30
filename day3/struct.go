package main

import "fmt"

type Customer struct {
  Name, Address string
  Age int
  
}

func main(){
  var farhan Customer
  farhan.Name = "farhan"
  farhan.Address ="perum tti"
  farhan.Age = 22

  fmt.Println(farhan)
  fmt.Println(farhan.Nama)
  fmt.Println(farhan.Address)
  fmt.Println(farhan.Age)
 
  //another way to build struct

  farhan := Customer {
    Name : "Farhan",
    Address : "Perum tti",
    Age : 22,
    }
}
