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
	hasil, err := pembagi(10, 0)

	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("terjadi error: ", err.Error())
	}
}
