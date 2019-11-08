package sql_injection

import (
	"database/sql"
	"fmt"
)

type UserDetails struct {
	Id         string
	CardNumber string
}

func GetCardNumber(db *sql.DB, userID string) (resp string, err error) {
	query := `SELECT CARD_NUMBER FROM USER_DETAILS WHERE USER_ID = ` + userID
	row := db.QueryRow(query)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	case nil:
		return resp, err
	default:
		return resp, err
	}
	return
}
func GetCardNumberSecure(db *sql.DB, userID string) (resp string, err error) {
	stmt, err := db.Prepare(`SELECT CARD_NUMBER FROM USER_DETAILS WHERE USER_ID = ?`)
	if err != nil {
		return resp, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(userID)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	case nil:
		return resp, err
	default:
		return resp, err
	}
	return
}
