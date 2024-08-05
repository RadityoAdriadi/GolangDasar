package main

import "fmt"

func main() {
	salary := 7000000

	if salary != 650000 {
		fmt.Println("Kerja Rodi")
	} else {
		fmt.Println("Budak Koorporate")
	}

	working_hour_perday := 8
	POV_PT := 10

	if working_hour_perday != POV_PT {
		fmt.Println("Kurang Loyal")
	} else if working_hour_perday == POV_PT {
		fmt.Println("Sangat Loyal")
	} else {
		fmt.Println("Work Life Ballance")
	}

	gaji := 2000000
	umk := 5100000
	ump := 1700000

	if gaji < umk {
		fmt.Println("Gila sih kamu di kerjaiin doang itu masa cuman Rp.", gaji)
		if gaji > ump {
			fmt.Println("Alhamdulillah, bersyukur dulu aja gaji kamu Rp.", gaji, ", masih lebih gede dari anak bu siti cuman Rp.", ump)
		}
	}

}
