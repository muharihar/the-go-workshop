package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	db, err := sql.Open("postgres", "user=postgres password=admin host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(2, "second")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The value was successfully inserted!")
	}
	db.Close()
}
