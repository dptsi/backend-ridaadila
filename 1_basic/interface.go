package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHey(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{
		Name: "rida",
	}

	var cat Animal = Animal{
		Name: "cat",
	}

	sayHey(person)
	sayHey(cat)
}
