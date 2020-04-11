package base

import "os"

const env = "development"
const dbpath = "data/geo-velocity.sqlite3"
const encIV = "28472B4B6250655368566D597133743677397A24432646294A404D635166546A"
const encKey256 = "703273357638792F423F4528472B4B6250655368566D597133743677397A2443"
const sslEnable = "false"
const serverPort = "8080"
const skipSSLVerify = "false"
const host = "localhost"

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
func EncKey256() string {
	val := os.Getenv("encryption_key")

	if val == "" {
		val = encKey256
	}

	return val
}

func EncIV() string {
	val := os.Getenv("encryption_iv")

	if val == "" {
		val = encIV
	}

	return val
}

func EnableSSL() string {
	val := os.Getenv("enable_ssl")

	if val == "" {
		val = sslEnable
	}

	return val
}

func ServerPort() string {
	val := os.Getenv("server_port")

	if val == "" {
		val = serverPort
	}

	return val
}

// Do not do this in production
func SkipSSLVerify() string {
	val := os.Getenv("skip_ssl_verify")

	if val == "" {
		val = skipSSLVerify
	}

	return val
}
