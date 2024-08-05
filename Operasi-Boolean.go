package main

import (
	"fmt"
	"strconv"
)

func main() {
	var NilaiGaji = 2500000
	var BiayaHidup = 1500000

	var THP bool = NilaiGaji >= 2500000
	var biaya bool = BiayaHidup >= 1500000

	var cukup1 bool = THP && biaya
	var cukup2 bool = THP || biaya
	var cukup3 bool = !biaya
	var cukup4 bool = !THP
	var cukup5 bool = !THP || !biaya
	var cukup6 bool = !THP && !biaya
	fmt.Println(cukup1)
	fmt.Println(cukup2)
	fmt.Println(cukup3)
	fmt.Println(cukup4)
	fmt.Println(cukup5)
	fmt.Println(cukup6)

	ValueGaji := 5000000
	ValuePengeluaran := 5500000

	monthlyexpense := ValueGaji > ValuePengeluaran
	result2 := strconv.FormatBool(monthlyexpense)
	fmt.Println(result2)
	fmt.Printf(`Boros lu ah`, result2)

}
