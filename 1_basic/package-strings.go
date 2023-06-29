package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Rida Adila", "Rida"))
	fmt.Println(strings.ToUpper("rida adila"))
	fmt.Println(strings.ToLower("RIDA ADILA"))
	fmt.Println(strings.Trim("   rida    ", " "))
	fmt.Println(strings.ReplaceAll("rida tes adila", "tes", "with"))
}
