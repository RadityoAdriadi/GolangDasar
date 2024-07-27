package main

import "fmt"

func main() {
	const Religion string = "Islam"
	const Partner = "Dewi"
	const Salary = 5000

	fmt.Println(Religion)
	fmt.Println("My Beloved One =", Partner)
	fmt.Println("Gaji =", Salary, "USD")

	const (
		NameOfBirth string = "Ito"
		NameOfAdult        = "Radit"
		weight             = 115
	)

	fmt.Println(NameOfBirth)
	fmt.Println(NameOfAdult)
	fmt.Println("Berat Badan =", weight)
}
