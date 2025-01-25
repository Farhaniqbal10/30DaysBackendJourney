package main

import "fmt"

func main(){
	var num1, num2 float64
	var operator string

	fmt.Println("kalkulator sederhana")
	fmt.Println("masukan operator (+,-,*,/)")
	fmt.Scanln(&operator)

	fmt.Println("masukan angka pertama")
	fmt.Scanln(&num1)

	fmt.Println("masukan angka kedua")
	fmt.Scanln(&num2)

	switch operator {
	case "+" :
		fmt.Printf("hasil akhir : %.2f\n", num1+num2)
	case "-" : 
		fmt.Printf("hasil akhir : %.2f\n", num1-num2)
	case "*": 
		fmt.Printf("hasil akhir : %.2f\n", num1*num2)
	case "/" :
		if num2 != 0 {
		fmt.Printf("hasil akhir : %.2f\n", num1/num2)	
		}else{
			fmt.Println("error! kamu tidak bisa membagi dengan 0")
		}
		default :
		fmt.Println("error! operator yang kamu masukan salah")
	}
}