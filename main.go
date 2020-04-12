package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	connectDB()
}

func connectDB() {
	var err error
	connectionString := "postgres://zecolkej:KwZAmg...@john.db.elephantsql.com:5432/zecolkej"
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect DB Success\n", conn)
}
