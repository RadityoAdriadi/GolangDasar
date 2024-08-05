package main

import (
	"fmt"
)

func main() {
	// myslice1 := []int{}
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))
	// fmt.Println(myslice1)

	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))
	// fmt.Println(myslice2)

	DataUser := []string{
		"227",
		"Sudirman",
		"Surabaya",
		"01/07/2024",
		"Sudah menikah",
		"Dewi Tiaman Putri",
		"Pegawai Swasta",
	}
	fmt.Println("Data Pegawai")
	fmt.Println("ID Pegawai = ", DataUser[0])
	fmt.Println("Nama Pegawai = ", DataUser[1])
	fmt.Println("Tempat Lahir = ", DataUser[2])
	fmt.Println("Tanggal Lahir = ", DataUser[3])
	fmt.Println("Status Pernikahan = ", DataUser[4])
	fmt.Println("Nama Pasangan = ", DataUser[5])
	fmt.Println("Pegawai Swasta = ", DataUser[6])
	// fmt.Println(len(DataUser))
	// fmt.Println(cap(DataUser))

}
