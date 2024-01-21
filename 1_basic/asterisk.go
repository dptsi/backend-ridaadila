package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{City: "Subang", Province: "Jawa Barat", Country: "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{City: "Surabaya", Province: "Jawa Timur", Country: "Indonesia"}

	fmt.Println(address2)
	fmt.Println(address1)
}
