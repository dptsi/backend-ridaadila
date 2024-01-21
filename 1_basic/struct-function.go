package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (person Customer) sayGoodbye(city string) {
	fmt.Println("goodbye ", person.Name, ", from ", city)
}

func main() {

	person := Customer{
		Name:    "rida",
		Address: "bandung",
		Age:     22,
	}

	person.sayGoodbye("new york")
}
