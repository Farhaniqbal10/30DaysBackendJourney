package main

import "fmt"

func faktorial(n int) int {
	result := 1 
	for i := 1; i <= n; i++{
		result *= i
	}
	return result
}

func main(){
	var num1, num2 float64
	var operator string
	var num int

	fmt.Println("kalkulator sederhana")
	fmt.Println("masukan operator atau pilih ! untuk faktorial (+,-,*,/,!) ")
	fmt.Scanln(&operator)
	
	if operator == "!" {
		fmt.Println("masukan angka yang ingin difaktorialkan :")
		fmt.Scanln(&num)

		if num < 0 {
			fmt.Println("masukan angka yang lebih dari 0")
		} else {
			result := faktorial(num)
			fmt.Printf("faktorial dari %d adalah %d\n", num, result)
		}
	} else {

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
}