package main 

import "fmt"

// Fungsi untuk menghitung luas persegi panjang

func luasPersegiPanjang(lebar float64, panjang float64){
	return lebar*panjang
}

// Fungsi untuk menampilkan salam

func sapa(nama string) string {
	return "halo" + nama + "selamat datang di program go"
}

func main(){

	var nama string = "farhan"
	umur := 23
	var tinggi float64 = 170.5
	var aktif bool = true


	fmt.Println("nama :", nama)
	fmt.Println("umur :", umur)
	fmt.Println("tinggi :", tinggi)
	fmt.Println("status aktif :", aktif)

	panjang := 10.5
	lebar := 5.2
	luas := luasPersegiPanjang(lebar, panjang)
	fmt.Println("luas persegi panjang :", luas)

	fmt.Println(sapa)


	

}
