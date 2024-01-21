package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountry(address *Address) {
	address.Country = "Singapore"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", "Indonesia"}

	ChangeCountry(&address)

	fmt.Println(address)
}
