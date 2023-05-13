package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "rida"
	nama[1] = "adila"
	nama[2] = "tes"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var nilai = [3]int{
		1, 2, 3,
	}

	fmt.Println(nilai)

	fmt.Println(len(nilai))

	var kosong [11]string

	fmt.Println(len(kosong))
}
