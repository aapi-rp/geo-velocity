package utils

import (
	"math"
	"time"
)

func getCurrentEpochTime() time.Time {
	now := time.Now()
	secs := now.Unix()
	return time.Unix(secs, 0)
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

const radius = 6371 // Earth's avarage radius in kilometers

func deg2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Distance(origin Coordinates, destination Coordinates) float64 {
	dgLat := deg2radians(destination.Latitude - origin.Latitude)
	dgLong := deg2radians(destination.Longitude - origin.Longitude)
	a := math.Sin(dgLat/2)*math.Sin(dgLong/2) + math.Cos(deg2radians(origin.Latitude))*math.Cos(deg2radians(destination.Latitude))*math.Sin(dgLong/2)*math.Sin(dgLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return d
}
