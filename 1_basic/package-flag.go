package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "default host", "Put your host")
	var user *string = flag.String("user", "default user", "Put your username")

	flag.Parse()

	fmt.Println("Hostname: ", *host)
	fmt.Println("Username: ", *user)
}
