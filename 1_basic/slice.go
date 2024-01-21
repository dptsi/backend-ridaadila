package main

import "fmt"

func main() {
	var array = [...]string{
		"rida", "test", "coba", "lagi",
	}

	slice1 := array[1:3]
	fmt.Println(slice1)

	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	dayslice1 := days[5:]

	fmt.Println(dayslice1)

	dayslice1[0] = "sabtu baru"
	dayslice1[1] = "minggu baru"

	fmt.Println(days)

	dayslice2 := append(dayslice1, "new days")

	fmt.Println(dayslice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "new slice 0"
	newSlice[1] = "new slice 1"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "append data")
	newSlice2[0] = "berubah"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
