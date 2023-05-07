package main

import "fmt"

func main() {
	type flag bool
	type uuid string

	var status flag = false
	var idUser uuid = "1231-232137871-12312"

	fmt.Println(status)
	fmt.Println(idUser)
}
