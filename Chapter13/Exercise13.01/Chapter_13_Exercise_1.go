package main

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"

func main(){
var prop string

db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

if err != nil {
  panic(err)
}else{
	fmt.Println("The connection to the DB was successfully initialized!")
}

TableCreate := `	
CREATE TABLE Number
(
    Number integer NOT NULL,
    Property text COLLATE pg_catalog."default" NOT NULL
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE Number
    OWNER to postgres;
`
_,err = db.Exec(TableCreate)
if err != nil {
	panic(err)
} else{
	fmt.Println("The table called Messages was successfully created!")
}

insert, insertErr := db.Prepare("INSERT INTO Number VALUES($1,$2)")
if insertErr != nil{
	panic(insertErr)
}

for i := 0; i < 100; i++ {
	if i % 2 == 0{
   		prop = "Even"
	}else{
    	prop = "Odd"
    }
	_, err = insert.Exec(i,prop)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The number:",i,"is:",prop)
	}
}
insert.Close()

fmt.Println("The numbers are ready.")
db.Close()

}