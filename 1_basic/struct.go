package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var rida Customer

	rida.Name = "Rida Adila"
	rida.Address = "Surabaya"
	rida.Age = 23

	fmt.Println(rida)
	fmt.Println(rida.Age)

	person := Customer{
		Name:    "Christian",
		Address: "Jakarta",
		Age:     37,
	}

	fmt.Println(person)
}
