package main

import "fmt"

func getInfo(nama string) string {
	return "Hi " + nama
}

func main() {
	fmt.Println(getInfo("rida"))
}
