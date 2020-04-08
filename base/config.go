package base

import "os"

const env = "development"
const dbpath = "/Users/rphillips/Documents/github.com/aapi-rp/geo-velocity/db/geo-velocity.sqlite3"

func GetEnv() string {
	val := os.Getenv("env")

	if val == "" {
		val = env
	}

	return val
}

func DBPath() string {
	val := os.Getenv("sqlite3_db_path")

	if val == "" {
		val = dbpath
	}

	return val
}
