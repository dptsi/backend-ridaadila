package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rida")

	if result != "Hello Rida" {
		panic("Result is not hello rida. failed")
	}
}
