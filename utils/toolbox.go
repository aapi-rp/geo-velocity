package utils

import (
	"time"
)

func getCurrentEpochTime() time.Time {
	now := time.Now()
	secs := now.Unix()
	return time.Unix(secs, 0)
}
