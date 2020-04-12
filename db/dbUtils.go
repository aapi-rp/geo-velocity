package db

import (
	"database/sql"
	"github.com/aapi-rp/geo-velocity/config"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/model_struct"
	_ "github.com/mattn/go-sqlite3"
)

var SqliteConn *sql.DB

/*
   Method Description: Creates the initial database.
*/
func InitData() (*sql.DB, error) {

	DbConn, err := sql.Open("sqlite3", config.DBPath()+"?cache=shared&mode=rwc")

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

/*
    Method Description: Generic method to insert data into the events database, this keeps developers always pushing to the database
    in the same way, sqlite is an interesting database, postgres or mysql would have been a better sampling of my database skills
    we could have just used an inmem mysql and spun it up on the same docker auto.
	Parameter: query
    Parameter Description: The insert query you would like to submit
	Parameter: values
    Parameter Description: Standard values for the insert
    Returns any error that may come from the database
*/
func InsertDBRow(query string, values ...interface{}) error {

	DbConn, err := sql.Open("sqlite3", config.DBPath()+"?cache=shared&mode=rwc")

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

/*
    Method Description: Method to see if row exists
	Parameter: query
    Parameter Description: The insert query you would like to submit
	Parameter: args
    Parameter Description: Standard values for the query
    Returns true or false if the row exists, or any error that may come from the database
*/
func SelectDBRowExists(query string, args ...interface{}) (bool, error) {

	DbConn, err := sql.Open("sqlite3", config.DBPath()+"?cache=shared&mode=rwc")

	if err != nil {
		return false, err
	}

	err = DbConn.Ping()

	if err != nil {
		return false, err
	}

	rows, err := DbConn.Query(query, args...)

	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}

	rows.Close()

	return false, nil
}

/*
    Method Description: Method to see if row exists
	Parameter: query
    Parameter Description: The insert query you would like to submit
	Parameter: args
    Parameter Description: Standard values for the query
    Parameter: currentIPAddr
    Parameter Description: ip for logging
	Parameter: currentUser
    Parameter Description: user for logging
    Returns true or false if the row exists, geolocation data of the queried object and any error that may come from the database
*/
func SelectDBRows(query string, currentIPAddr string, currentUser string, args ...interface{}) (model_struct.GeoData, bool, error) {

	subGeoData := model_struct.GeoData{}
	hasrows := false
	DbConn, err := sql.Open("sqlite3", config.DBPath()+"?cache=shared&mode=rwc")

	if err != nil {
		return subGeoData, false, err
	}

	err = DbConn.Ping()

	if err != nil {
		return subGeoData, false, err
	}

	subRows, err := DbConn.Query(query, args...)

	for subRows.Next() {
		err = subRows.Scan(&subGeoData.IP_ADDRESS, &subGeoData.LAT, &subGeoData.LONG, &subGeoData.RADIUS, &subGeoData.LOGIN_TIME)
		if err != nil {
			logger.Error("Something happened while scanning the rows for: ", currentIPAddr, currentUser, " for query: ", query)
			return subGeoData, false, err
		}
		hasrows = true
	}

	subRows.Close()

	return subGeoData, hasrows, nil

}
