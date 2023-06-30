package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	utc := time.Date(2023, 12, 10, 10, 10, 10, 10, time.UTC)

	fmt.Println(utc)

	parse, _ := time.Parse("2006-01-02", "2023-06-10")
	fmt.Println(parse)
}
