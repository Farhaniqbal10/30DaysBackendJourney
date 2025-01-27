package main

import "fmt"

// Fungsi untuk mencari bilangan genap dalam array
func cariBilanganGenap(arr []int) []int {
	var bilanganGenap []int
	for _, nilai := range arr {
		if nilai%2 == 0 {
			bilanganGenap = append(bilanganGenap, nilai)
		}
	}
	return bilanganGenap
}

func main() {
	// Array angka
	arrayAngka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Menampilkan array asli
	fmt.Println("Array asli:", arrayAngka)

	// Memanggil fungsi cariBilanganGenap
	bilanganGenap := cariBilanganGenap(arrayAngka)

	// Menampilkan bilangan genap
	fmt.Println("Bilangan genap dalam array:", bilanganGenap)
}
