package main

import "fmt"

func getInfo(nama string, umur int) (string, int) {
	return nama, umur
}

func getStatus() (bool, string) {
	return true, "sukses"
}

func main() {
	name, age := getInfo("rida adila", 23)

	fmt.Println("Nama : ", name)
	fmt.Println("Umur: ", age)

	_, namaStatus := getStatus()
	fmt.Println(namaStatus)
}
