package main

import "fmt"

func faktorial(n int) int {
	result := 1 
	for i := 1; i <= n; i++{
		result *= i
	}
	return result
}

func main() {
	var num int

	// Input angka
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&num)

	// Cek apakah angka yang dimasukkan valid
	if num < 0 {
		fmt.Println("Angka tidak boleh negatif")
	} else {
		// Hitung dan tampilkan hasil faktorial
		result := faktorial(num)
		fmt.Printf("Faktorial dari %d adalah %d\n", num, result)
	}
}