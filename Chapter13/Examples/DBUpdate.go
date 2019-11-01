package main

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"
func main(){

	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	
	if err != nil {
	  panic(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	
UpdateStatement :=`
	UPDATE test 
	SET name = $1 
	WHERE id = $2
	`
	
	UpdateResult, UpdateResultErr := db.Exec(UpdateStatement,"well",2)
	if UpdateResultErr != nil {
		panic(UpdateResultErr)
	}
	UpdatedRecords, UpdatedRecordsErr := UpdateResult.RowsAffected()
	if UpdatedRecordsErr != nil {
	  panic(UpdatedRecordsErr)
	}
	
	fmt.Println("Number of records updated:",UpdatedRecords)
	
	db.Close()
	}
		