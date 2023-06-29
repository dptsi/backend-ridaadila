package database

import "fmt"

var connection string

func init() {
	fmt.Println("init dipanggil")
	connection = "SQL Server"
}

func GetConnection() string {
	return connection
}
