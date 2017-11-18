package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 32768
	user     = "postgres"
	password = "qwe123"
	dbname   = "messeinfor_kownledge"
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	checkErr(err)
}
