package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "rida" {
		return &notFoundError{"id tdk ditemukan"}
	}

	return nil
}

func main() {
	err := SaveData("rida", nil)

	if err != nil {

		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error: ", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error: ", notFoundError.Error())
		} else {
			fmt.Println("Unknown error")
		}

	} else {
		fmt.Println("sukses")
	}
}
