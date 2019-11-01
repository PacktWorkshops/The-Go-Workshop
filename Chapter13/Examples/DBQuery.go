package main

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"

func main(){

	var id int
	var name string
	
	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	
	if err != nil {
	  panic(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err)
	}
	
	
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	
	rows.Close()
	db.Close()
	}
	
		