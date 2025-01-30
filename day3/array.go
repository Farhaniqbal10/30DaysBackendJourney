package main

import "fmt"

func main(){
  var names [3]string
  names[0] = "farhan"
  names[1] = "maulana"
  names[2] = "iqbal"

  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  var values = [3]int {
    15,
    12,
    11,
    }
  fmt.Println(values)

  var values2 = [...]int {
    10,
    15,
    25,
    20,
    35,
    }
  fmt.Println(values2)
  }
