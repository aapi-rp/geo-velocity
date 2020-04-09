package db

import (
	"database/sql"
	"github.com/aapi-rp/geo-velocity/base"
	_ "github.com/mattn/go-sqlite3"
)

var SqliteConn *sql.DB

func InitData() (*sql.DB, error) {

	DbConn, err := sql.Open("sqlite3", base.DBPath())
	if err != nil {
		return nil, err
	}

	ct, err := DbConn.Prepare(createGVTable)

	if err != nil {
		return nil, err
	}

	ct.Exec()

	err = DbConn.Ping()

	if err != nil {
		return nil, err
	}

	SqliteConn = DbConn

	return DbConn, nil
}

func InsertDBRow(query string, values ...interface{}) error {

	DbConn, err := sql.Open("sqlite3", base.DBPath())
	if err != nil {
		return err
	}

	err = DbConn.Ping()

	if err != nil {
		return err
	}

	insert, err := DbConn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = insert.Exec(values...)

	if err != nil {
		return err
	}

	return nil
}
