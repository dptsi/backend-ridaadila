package main

import "fmt"

func main() {

	num := 1

	for num <= 10 {
		fmt.Println("perulangan ke: ", num)
		num++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("perulangan dgn init dan post statemen ke: ", i)
	}

	slice := []string{"rida", "adilaa", "tes", "123"}

	for _, value := range slice {
		fmt.Println(value)
	}

	maps := map[string]string{
		"nama":   "rida",
		"alamat": "surabaya",
		"umur":   "22",
	}

	for key, val := range maps {
		fmt.Println(key, val)
	}
}
