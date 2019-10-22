package exercise1
import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)
var db *sql.DB
type UserDetails struct {
	Id         string
	CardNumber string
	Address    string
}
var testData = []*UserDetails{
	{
		Id:         "1",
		CardNumber: "1234",
		Address:    "1 Cook Creet",
	},
	{
		Id:         "2",
		CardNumber: "5678",
		Address:    "2 Chi Creek",
	},
}
func getConnection() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "test.DB")
	if err != nil {
		return nil, fmt.Errorf("could not open db connection %v", err)
	}
	return conn, nil
}
func initializeDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USER_DETAILS (USER_ID TEXT, PHONE TEXT, ADDRESS TEXT)`)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(`INSERT INTO USER_DETAILS (USER_ID, PHONE, ADDRESS) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	for _, user := range testData {
		_, err := stmt.Exec(user.Id, user.CardNumber, user.Address)
		if err != nil {
			return err
		}
	}
	return nil
}
func tearDownDB(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE USER_DETAILS")
	if err != nil {
		return err
	}
	return nil
}
func TestMain(m *testing.M) {
	var err error
	db, err = getConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tearDownDB(db)
	if m.Run() != 0 {
		fmt.Println("error running tests")
		os.Exit(1)
	}
}
func TestUpdatePhoneSecure(t *testing.T) {
	var tests = []struct {
		ID    string
		Phone string
		err   string
	}{
		{
			ID:    "1",
			Phone: "1234",
			err:   "",
		},
		{
			ID:    "'2' OR USER_ID=='1'",
			Phone: "5678",
			err:   "no row affected",
		},
	}
	for _, test := range tests {
		err := UpdatePhoneSecure(db, test.ID, test.Phone)
		if err != nil {
			assert.EqualError(t, err, test.err)
		}
	}
}
