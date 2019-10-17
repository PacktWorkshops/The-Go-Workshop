package main

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"
import "math/big"

func main(){
var number int64
var prop string
var primeSum int64
var newNumber int64

db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
if err != nil {
  panic(err)
}else{
	fmt.Println("The connection to the DB was successfully initialized!")
}

AllTheNumbers := "SELECT * FROM Number"
Numbers, err := db.Prepare(AllTheNumbers)
if err != nil {
	panic(err)
} 

primeSum = 0
result, err := Numbers.Query()
fmt.Println("The list of prime numbers:")
for result.Next(){
	    err = result.Scan(&number, &prop)
	    if err != nil{
	    	panic(err)
	    }
	    if big.NewInt(number).ProbablyPrime(0) {
		    primeSum += number
		    fmt.Print(" ",number)
		}
		
	}

Numbers.Close()
fmt.Println("\nThe total sum of prime numbers in this range is:",primeSum)


Remove := "DELETE FROM Number WHERE Property=$1"
removeResult, err := db.Exec(Remove,"Even")
if err != nil {
	panic(err)
}
ModifiedRecords, err := removeResult.RowsAffected()
fmt.Println("The number of rows removed:",ModifiedRecords)
fmt.Println("Updating numbers...")

Update := "UPDATE Number SET Number=$1 WHERE Number=$2 AND Property=$3"
AllTheNumbers = "SELECT * FROM Number"
Numbers, err = db.Prepare(AllTheNumbers)
if err != nil {
	panic(err)
} 
result, err = Numbers.Query()
for result.Next(){
	    err = result.Scan(&number, &prop)
	    if err != nil{
	    	panic(err)
	    }
		newNumber = number + primeSum
		_, err = db.Exec(Update,newNumber,number,prop)
		if err != nil {
			panic(err)
		} 
	}
Numbers.Close()
fmt.Println("The execution is now complete...")
db.Close()

}