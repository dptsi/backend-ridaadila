package main

import "fmt"

type Filter func(string) string

func sayHello(nama string, filter Filter) {
	fmt.Println("Hello ", filter(nama))
}

func filterNama(nama string) string {
	if nama == "Anjing" {
		return "..."
	} else {
		return nama
	}
}

func main() {
	sayHello("rida", filterNama)
}
