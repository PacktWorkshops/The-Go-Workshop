package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "database/sql"
import _ "github.com/lib/pq"

type Messages struct {
    UserID int
    Message string
}

func main(){

var toLookFor string
var message string
var email string
var name string

reader := bufio.NewReader(os.Stdin)

messages := []Messages{
	{1,"Hi Florian, when are you coming home?"},
	{1,"Can you send some cash?"},
	{2,"Hi can you bring some bread and milk?"},
	{7,"Well..."},
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
CREATE TABLE public.messages
(
    UserID integer NOT NULL,
    Message character varying(280) COLLATE pg_catalog."default" NOT NULL
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.messages
    OWNER to postgres;
`

_,err = db.Exec(TableCreate)
if err != nil {
	panic(err)
} else{
	fmt.Println("The table called Messages was successfully created!")
}


insertMessages, insertErr := db.Prepare("INSERT INTO messages VALUES($1,$2)")
if insertErr != nil{
	panic(insertErr)
}

for _, u := range messages{
	_, err = insertMessages.Exec(u.UserID,u.Message)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The UserID:",u.UserID,"with message:",u.Message,"was successfully added!")
	}
}
insertMessages.Close()

fmt.Print("Give me the user's name: ")

toLookFor, err = reader.ReadString('\n')
toLookFor = strings.TrimRight(toLookFor, "\r\n")
if err != nil{
	panic(err)
} else {
	fmt.Println("Looking for all the messages of user with name:",toLookFor,"##")
}

UserMessages := "SELECT users.Name, users.Email, messages.Message FROM messages INNER JOIN users ON users.ID=messages.UserID WHERE users.Name LIKE $1"

usersMessages, err := db.Prepare(UserMessages)
if err != nil {
	panic(err)
} 

result, err := usersMessages.Query(toLookFor)

numberof := 0
for result.Next(){
   numberof++
}
if numberof == 0 {
   fmt.Println("The query returned nothing, no such user:",toLookFor)
}else{
	fmt.Println("There are a total of",numberof,"messages from the user:",toLookFor)

	result, err := usersMessages.Query(toLookFor)
	for result.Next(){
	    err = result.Scan(&name, &email, &message)
	    if err != nil{
	    	panic(err)
	    }
		fmt.Println("The user:",name,"with email:",email,"has sent the following message:",message)
	}
}
usersMessages.Close()
db.Close()
}

