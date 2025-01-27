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
	// Tipe data dan variabel

	var nama string = "farhan"  		//tipe data string
	umur := 23				//tipe data int
	var tinggi float64 = 170.5		//tipe data float
	var aktif bool = true			//tipe data boolean

	//menampillkan nilai variable
	fmt.Println("nama :", nama)
	fmt.Println("umur :", umur)
	fmt.Println("tinggi :", tinggi)
	fmt.Println("status aktif :", aktif)

	//proses perhitungan luas persegi panjang

	panjang := 10.5
	lebar := 5.2
	luas := luasPersegiPanjang(lebar, panjang)
	fmt.Println("luas persegi panjang :", luas)

	fmt.Println(sapa)


	

}
