package main

import "fmt"

func main() {
	var email1 = "raditya@gmail.com"
	var email2 = "raditya@gmail.com"

	var hasil bool = email1 == email2
	fmt.Println(hasil)

	Nilai1 := 100
	Nilai2 := 200
	var result bool = Nilai1 > Nilai2
	fmt.Println(result)

	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = value == 2
	fmt.Println(isEqual)

	var value2 = (((2+6)%3)*4 - 2) / 3
	var isEqual2 = value2 != 2
	fmt.Println(isEqual2)
}
