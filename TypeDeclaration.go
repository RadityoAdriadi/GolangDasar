package main

import "fmt"

func main() {
	type NPWP string
	type married bool

	var NPWP2 NPWP = "32410102121"
	var maritalstatus married = false

	fmt.Println(NPWP2)
	fmt.Println(maritalstatus)
}
