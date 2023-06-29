package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolval, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolval)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	value := strconv.FormatInt(1000, 10)
	fmt.Println(value)

}
