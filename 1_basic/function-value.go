package main

import "fmt"

func sayHello(nama string) string {
	return "Halo " + nama
}

func main() {
	halo := sayHello

	fmt.Println(halo("rida"))
}
