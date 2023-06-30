package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := UserSlice{
		{"Rida", 23},
		{"Tyas", 30},
		{"Michael", 28},
		{"Herra", 17},
	}

	sort.Sort(users)

	fmt.Println(users)
	fmt.Println(users[1].Name)
}
