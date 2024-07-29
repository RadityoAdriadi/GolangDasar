package main

import "fmt"

func main() {
	var hasil = 50 + 50*100
	fmt.Println(hasil)

	var sangu = 10000
	var cilok = 5000
	var ditabung = cilok - sangu
	fmt.Println(ditabung)

	//augmanted assigment
	var pempek = 5000
	pempek += 15000

	fmt.Println(pempek)

	var permen = 10000
	permen *= 3

	fmt.Println(permen)

	//unary operator
	var a = 10
	a++
	fmt.Println(a)

	var b = 10
	b--
	fmt.Println(b)

	var c = -100
	c--
	fmt.Println(c)

	var d = -2000
	fmt.Println(d)
}
