package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile("r([a-z])da")

	fmt.Println(regex.MatchString("tia"))
	fmt.Println(regex.MatchString("rida"))

	fmt.Println(regex.FindAllString("rida roda apa", 1))
}
