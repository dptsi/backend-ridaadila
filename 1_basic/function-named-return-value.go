package main

import "fmt"

func getFullName() (firstname, middlename, lastname string) {

	firstname = "rida"
	middlename = "test"
	lastname = "adila"

	return
}

func main() {

	_, _, lastname := getFullName()

	fmt.Println(lastname)
}
