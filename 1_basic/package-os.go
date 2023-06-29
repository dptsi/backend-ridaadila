package main

import (
	"fmt"
	"os"
)

func main() {
	osvalue := os.Args
	fmt.Println(osvalue)

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
