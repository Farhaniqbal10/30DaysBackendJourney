package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai pangkat
func hitungPangkat(base int, exponent int) int {
	// Basis rekursi: jika eksponen 0, hasilnya 1
	if exponent == 0 {
		return 1
	}
	// Rekursi: base * hasil dari (base ^ (exponent - 1))
	return base * hitungPangkat(base, exponent-1)
}

func main() {
	// Contoh penggunaan fungsi hitungPangkat
	base := 2
	exponent := 3

	hasil := hitungPangkat(base, exponent)
	fmt.Printf("%d^%d = %d\n", base, exponent, hasil)
}
