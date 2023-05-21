package main

import "fmt"

func getInfo(nama string, umur int) {
	fmt.Println("Nama: ", nama, "\nUmur: ", umur)
}

func main() {
	getInfo("rida", 22)
}
