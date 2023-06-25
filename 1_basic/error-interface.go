package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		errors := errors.New("pembagi tdk boleh nol")
		return 0, errors
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	fmt.Println(pembagi(10, 0))
}
