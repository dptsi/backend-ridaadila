package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[5:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	bulan[5] = "bulan ubah"
	fmt.Println(slice1)
	slice1[1] = "Juli ubah"
	fmt.Println(bulan)

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[4:6]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days)

	daysAppend := append(daySlice1, "Tes")
	daysAppend[0] = "Sabtu baru berubah"
	fmt.Println(daysAppend)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "index 1"
	newSlice[1] = "index 2"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)
}
