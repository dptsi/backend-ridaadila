package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToEngland(address *Address) {
	address.Country = "England"
}

func main() {
	var address1 Address = Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}
	var address2 *Address = &address1
	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Bandung"

	fmt.Println(address4)

	var alamat = Address{
		City:     "London",
		Province: "Ada",
		Country:  "",
	}
	fmt.Println(alamat)
	ChangeCountryToEngland(&alamat)
	fmt.Println(alamat)
}
