package main

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE x (name TEXT, age INT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO x VALUES ('Terry', 12), ('Marge', 14)")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT name, age FROM x")
	for rows.Next() {
		var name string
		var age int
		err = rows.Scan(&name, &age)
		if err != nil {
			panic(err)
		}

		fmt.Println(name, age)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
