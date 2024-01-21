package main

import "fmt"

func main() {

	person := map[string]string{
		"nama":   "rida",
		"alamat": "surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

	person["umur"] = "22"

	fmt.Println(person["umur"])

	buku := make(map[string]string)
	buku["judul"] = "buku bahasa inggris"
	buku["pengarang"] = "andi"

	fmt.Println(buku)

	delete(buku, "judul")

	fmt.Println(buku)
	fmt.Println(len(buku))
}
