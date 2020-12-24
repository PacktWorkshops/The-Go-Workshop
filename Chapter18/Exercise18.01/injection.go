package exercise1
import (
	"database/sql"
	"fmt"
	"strings"
)
func UpdatePhone(db *sql.DB, ID string, phone string) error {
	var builder strings.Builder
	builder.WriteString("UPDATE USER_DETAILS SET PHONE=")
	builder.WriteString(phone)
	builder.WriteString(" WHERE USER_ID=")
	builder.WriteString(ID)
	fmt.Printf("Running query: %s\n", builder.String())
	_, err := db.Exec(builder.String())
	if err != nil {
		return err
	}
	return nil
}
func UpdatePhoneSecure(db *sql.DB, ID string, phone string) error {
	stmt, err := db.Prepare(`UPDATE USER_DETAILS SET PHONE=? WHERE USER_ID=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(phone, ID)
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
		return fmt.Errorf("more than one row affected")
	}
	return nil
}
