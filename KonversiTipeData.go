package main

import "fmt"

func main() {
	var nilai32 int32 = 130
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Radityo"
	var i = name[3]
	var iString string = string(i)

	fmt.Println("Radityo"[3])
	fmt.Println(name)
	fmt.Println(iString)

	var status = "Jadi Yatim"
	var s = status[5]
	var ConvtoString = string(s)

	fmt.Println(status)
	fmt.Println("Jadi Yatim"[5])
	fmt.Println(ConvtoString)
}
