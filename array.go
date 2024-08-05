package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Radit"
	names[1] = "ito"
	names[2] = "dito"

	fmt.Println(names[1])
	fmt.Println(names[1])
	fmt.Println(names[0])

	var values = [3]int{
		96,
		97,
		98,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var data [10]int

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(data))

}
