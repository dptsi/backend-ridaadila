package main

import "fmt"

func apapun(value int) interface{} {

	if value == 1 {
		return 1
	} else if value == 2 {
		return true
	} else {
		return "ok"
	}
}

func main() {
	var result interface{} = apapun(10)

	fmt.Println(result)
}
