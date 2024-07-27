package main

import "fmt"

func main() {
	var nama string

	nama = ("Radityo Adriadi")
	fmt.Println(nama)

	nama = ("Radityo")
	fmt.Println(nama)

	var NamaAyah = "Supriyadhi"
	fmt.Println(NamaAyah)

	var Usia = 28
	fmt.Println(Usia)

	var JenisKelamin = true
	fmt.Println("Laki-Laki =", JenisKelamin)

	country := "Indonesia"
	fmt.Println("Warga Negara = ", country)

	country = "Korea Utara"
	fmt.Println("Menetap di =", country)

	var (
		NamaIbu       = "Sri Lestari Astuti"
		Maritalstatus = "True"
	)
	fmt.Println("Nama Ibu =", NamaIbu)
	fmt.Println("Sudah Menikah =", Maritalstatus)
}
