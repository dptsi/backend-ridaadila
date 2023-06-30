package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data ke " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(a any) {
		fmt.Println(a)
	})
}
