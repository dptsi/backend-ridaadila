package main

import "fmt"

func main() {

	nilai := 9

	switch nilai {
	case 2:
		fmt.Println("nilai tdk sesuai dgn pengecekan")
	case 100:
		fmt.Println("nilai sesuai dengan pengecekan")
	default:
		fmt.Println("default print")
	}

	length := len("rida adila")

	switch {
	case length == 7:
		fmt.Println("panjang string 7")
	case length == 9:
		fmt.Println("panjang string 9")
	default:
		fmt.Println("panjang string 10")
	}
}
