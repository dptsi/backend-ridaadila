package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) changeName() {
	person.Name = "Changed Name"
}

func main() {
	rida := Person{
		Name: "Rida",
	}

	rida.changeName()
	fmt.Println(rida)
}
