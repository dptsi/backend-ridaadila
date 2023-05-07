package main

import "fmt"

func main() {
	const alamat string = "surabaya"
	const nama = "rida"
	const angka int = 100

	fmt.Println(nama)
	fmt.Println(alamat)

	const (
		flag   int8 = 1
		status      = "ok"
	)

	fmt.Println(status)
}
