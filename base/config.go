package base

import "os"

const env = "development"

func GetEnv() string {
	val := os.Getenv("env")

	if val == "" {
		val = env
	}

	return val
}
