package main

import "fmt"

func sumAll(numbers ...int) int {

	total := 0
	for _, val := range numbers {
		total += val
	}

	return total
}

func main() {
	total := sumAll(10, 200, 30, 40)

	fmt.Println(total)
	slice := []int{10, 20, 20, 9}

	by_slice := sumAll(slice...)
	fmt.Println(by_slice)
}
