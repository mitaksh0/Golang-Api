package database

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func DbInit() error {
	var err error
	Db, err = sql.Open("mysql", "root@tcp(localhost:3306)/go")
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("database connected successfully...")

	return err
}
