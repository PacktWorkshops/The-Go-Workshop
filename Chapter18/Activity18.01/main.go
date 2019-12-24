package main
import (
	"crypto/sha512"
	"database/sql"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
)
type UserDetails struct {
	Id       string
	Password string
}
var testData = []*UserDetails{
	{
		Id:       "1",
		Password: "1234",
	},
	{
		Id:       "2",
		Password: "5678",
	},
}
func initializeDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USER_DETAILS (USER_ID TEXT, PASSWORD TEXT)`)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(`INSERT INTO USER_DETAILS (USER_ID, PASSWORD) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	for _, user := range testData {
		_, err := stmt.Exec(user.Id, user.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
func tearDownDB(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE USER_DETAILS")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func getConnection() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "test.DB")
	if err != nil {
		return nil, fmt.Errorf("could not open db connection %v", err)
	}
	return conn, nil
}
func UpdatePassword(db *sql.DB, Id string, Password string) error {
	query := `UPDATE USER_DETAILS SET PASSWORD=? WHERE USER_ID=?`
	cipher := sha512.Sum512([]byte(Password))
	fmt.Printf("storing encrypted password:\n%x\n", string(cipher[:]))
	result, err := db.Exec(query, string(cipher[:]), Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no row affected")
	}
	if rows > 1 {
		return fmt.Errorf("more than one row affected: %d", rows)
	}
	return nil
}
func GetPassword(db *sql.DB, userID string) (resp []byte, err error) {
	query := `SELECT PASSWORD FROM USER_DETAILS WHERE USER_ID = ?`
	row := db.QueryRow(query, userID)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	case nil:
		return resp, err
	default:
		return resp, err
	}
}
func main() {
	db, err := getConnection()
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
	err = UpdatePassword(db, "1", "NewPassword")
	if err != nil {
		fmt.Println("error updating password: ", err)
	}
	// Check password
	fmt.Println("retrieving hashed password from db")
	password, err := GetPassword(db, "1")
	if err != nil {
		fmt.Println("error retrieving password: ", err)
	}
	fmt.Println("checking password match")
	ciphertext := sha512.Sum512([]byte("NewPassword"))
	if string(ciphertext[:]) == string(password[:]) {
		fmt.Println("successful password match")
	}
}
