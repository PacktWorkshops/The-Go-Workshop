package sql_injection

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

var testData = []*UserDetails{
	{
		Id:         "1",
		CardNumber: "1234",
	},
	{
		Id:         "2",
		CardNumber: "5678",
	},
	{
		Id:         "1",
		CardNumber: "91011",
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USER_DETAILS (USER_ID TEXT, CARD_NUMBER TEXT)`)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`INSERT INTO USER_DETAILS (USER_ID, CARD_NUMBER) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	for _, user := range testData {
		_, err := stmt.Exec(user.Id, user.CardNumber)
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

func TestGetCardNumber(t *testing.T) {
	var testCase = []struct {
		userID string
	}{
		{
			userID: "1",
		},
		{
			userID: "'' OR USER_ID=='2'",
		},
	}
	for _, test := range testCase {
		resp, err := GetCardNumber(db, test.userID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("query response: %s", resp)
	}
}

func TestGetCardNumberSecure(t *testing.T) {
	var testCase = []struct {
		userID string
		resp   string
		err    string
	}{
		{
			userID: "1",
			resp:   "1234",
			err:    "",
		},
		{
			userID: "'' OR USER_ID=='2",
			resp:   "",
			err:    "no rows returned",
		},
	}
	for _, test := range testCase {
		resp, err := GetCardNumberSecure(db, test.userID)
		assert.Equal(t, resp, test.resp, "response should be same")
		if err != nil {
			assert.EqualError(t, err, test.err)
		} else {
			t.Logf("query response: %s", resp)
		}
	}
}
