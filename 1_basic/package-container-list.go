package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("rida")
	data.PushBack("adila")
	data.PushBack("end")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
