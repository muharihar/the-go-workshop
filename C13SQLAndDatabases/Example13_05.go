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

	UpdateStatement := `
	UPDATE test
	SET name = $1
	WHERE id = $2`

	UpdateResult, UpdateResulErr := db.Exec(UpdateStatement, "well", 2)
	if UpdateResulErr != nil {
		panic(UpdateResulErr)
	}
	UpdateRecords, UpdateRecordsErr := UpdateResult.RowsAffected()
	if UpdateRecordsErr != nil {
		panic(UpdateRecordsErr)
	}

	fmt.Println("Number of records updated: ", UpdateRecords)
	db.Close()
}
