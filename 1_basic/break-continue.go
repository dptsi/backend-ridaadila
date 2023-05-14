package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 6 {
			continue
		}

		if i == 9 {
			break
		}

		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
