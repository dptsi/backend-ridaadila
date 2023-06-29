package main

import "fmt"

type Address struct {
	City, Province, Country string
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
}
