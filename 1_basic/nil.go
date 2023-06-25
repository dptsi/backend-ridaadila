package main

import "fmt"

func NewMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
}

func main() {

	person := NewMap("")
	fmt.Println(person)
}
