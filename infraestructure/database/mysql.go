package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB() (*sql.DB,error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hexagonal")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}