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
DBCreate := `
	CREATE TABLE public.test
	(
		id integer,
		name character varying COLLATE pg_catalog."default"
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE public.test
		OWNER to postgres;
	`
	_,err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	} else{
		fmt.Println("The table was successfully created!")
	}
	db.Close()
}	
