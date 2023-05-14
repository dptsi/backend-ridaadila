package main

import "fmt"

func main() {
	nama := "rida"

	if nama == "rida" {
		fmt.Println("nama sdh sesuai dgn pengecekan")
	} else {
		fmt.Println("nama tidak sesuai dgn pengecekan")
	}

	if angka := 2; angka == 1 {
		fmt.Println("angka sudah benar")
	} else {
		fmt.Println("angka belum benar")
	}
}
