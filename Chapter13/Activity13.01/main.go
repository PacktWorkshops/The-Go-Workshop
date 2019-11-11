package main
import "fmt"
import "database/sql"
import _ "github.com/lib/pq"
type Users struct {
    id int
    name string
    email string
}
func main(){
users := []Users{
	{1,"Szabo Daniel","daniel@packt.com"},
	{2,"Szabo Florian","florian@packt.com"},
}
db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
if err != nil {
  panic(err)
}else{
	fmt.Println("The connection to the DB was successfully initialized!")
}
connectivity := db.Ping()
if connectivity != nil{
	panic(connectivity)
}else{
	fmt.Println("Good to go!")
}
TableCreate := `
CREATE TABLE users
(
    ID integer NOT NULL,
    Name text COLLATE pg_catalog."default" NOT NULL,
    Email text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Users_pkey" PRIMARY KEY (ID)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;
ALTER TABLE users
    OWNER to postgres;
`
_,err = db.Exec(TableCreate)
if err != nil {
	panic(err)
} else{
	fmt.Println("The table called Users was successfully created!")
}

insert, insertErr := db.Prepare("INSERT INTO users VALUES($1,$2,$3)")
if insertErr != nil{
	panic(insertErr)
}
for _, u := range users{
	_, err = insert.Exec(u.id,u.name,u.email)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The user with name:",u.name,"and email:",u.email,"was successfully added!")
	}
}
insert.Close()
update, updateErr := db.Prepare("UPDATE users SET Email=$1 WHERE ID=$2")
if updateErr != nil{
	panic(updateErr)
}
_, err = update.Exec("user@packt.com",1)
if err != nil{
	panic(err)
} else{
	fmt.Println("The user's email address was successfully updated!")
}
update.Close()
remove, removeErr := db.Prepare("DELETE FROM users WHERE ID=$1")
if removeErr != nil{
	panic(removeErr)
}
_,err = remove.Exec(2)
if err != nil{
	panic(err)
}else{
	fmt.Println("The second user was successfully removed!")
}
remove.Close()
db.Close()
}
