package main

import (
	"demo/dbfile"
	"demo/httpfile"
)

func main() {
	// Please enter the local database password and name
	dsn := "root:544688504l@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	dc, err := dbfile.NewClient(dsn) //dc: dbClient
	if err != nil {
		panic(err)
	}

	server1 := httpfile.NewServer(dc)
	server1.Run(":8080")

}
