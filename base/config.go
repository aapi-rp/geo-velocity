package base

import "os"

const env = "development"
const dbpath = "/Users/rphillips/Documents/github.com/aapi-rp/geo-velocity/db/geo-velocity.sqlite3"
const encKey = "0E&@w85hetEO7rl2"

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

// This should be added as a Kubernetes secret key header so its protected.
// Default should not be used in production
func EncKey() string {
	val := os.Getenv("encryption_key")

	if val == "" {
		val = encKey
	}

	return val
}
