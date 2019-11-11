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
	DeleteStatement :=`
	DELETE FROM test 
	WHERE id = $1
	`		
	DeleteResult, DeleteResultErr := db.Exec(DeleteStatement,2)
if DeleteResultErr != nil {
	panic(DeleteResultErr)
}
DeletedRecords, DeletedRecordsErr := DeleteResult.RowsAffected()
if DeletedRecordsErr != nil {
  panic(DeletedRecordsErr)
}
fmt.Println("Number of records deleted:",DeletedRecords)	
	db.Close()
	}
