package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai dijalankan")
	error := recover()
	fmt.Println("recover", error)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("TERJADI ERROR")
	}

	fmt.Println("Aplikasi running")
}

func main() {
	runApp(true)
	fmt.Println("masuk")
}
