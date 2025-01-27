package main

import "fmt"

func main () {

	// Contoh Struktur Kontrol: IF
	nilai := 85
	fmt.Println("Nilai Anda:", nilai)
	if nilai >= 90 {
		fmt.Println("Grade: A")
	} else if nilai >= 75 {
		fmt.Println("Grade: B")
	} else if nilai >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// Contoh Struktur Kontrol: FOR
	fmt.Println("\nMenampilkan angka 1 sampai 5 menggunakan for:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka:", i)
	}


	fmt.Println("\nMenampilkan bilangan genap antara 1 sampai 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//contoh struktur kontrol switch case
	hari := 3
	fmt.Print("\nHari ke-", hari, " adalah ")
	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak valid")
	}
}
