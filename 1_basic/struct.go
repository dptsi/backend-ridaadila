package main

import "fmt"

type Customer struct {
	Nama, Alamat string
	Umur         int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("hello", name, "my name is", customer.Nama)
}

func main() {

	var person1 Customer

	person1.Nama = "rida"
	person1.Alamat = "surabaya"
	person1.Umur = 23

	fmt.Println(person1)
	fmt.Println(person1.Nama)

	var person2 Customer = Customer{
		Nama: "Michael",
		Umur: 40,
	}

	fmt.Println(person2)
	fmt.Println("Alamat person2:", person2.Alamat)

	person3 := Customer{
		Nama:   "Hiroshi",
		Alamat: "Tokyo",
		Umur:   30,
	}

	fmt.Println(person3)

	person1.sayHi("james")
}
