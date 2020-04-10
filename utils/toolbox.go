package utils

import (
	"github.com/aapi-rp/geo-velocity/model_struct"
	"github.com/oschwald/geoip2-golang"
	"log"
	"math"
	"net"
	"time"
)

func getCurrentEpochTime() time.Time {
	now := time.Now()
	secs := now.Unix()
	return time.Unix(secs, 0)
}

func GetGeoDataFromIP(ipaddr string) (model_struct.GeoData, error) {

	Geo := model_struct.GeoData{}

	db, err := geoip2.Open("db/GeoLite2-City.mmdb")
	if err != nil {
		log.Println(err)
		return Geo, err
	}
	defer db.Close()

	//67.181.148.192
	ip := net.ParseIP(ipaddr)
	record, err := db.City(ip)
	if err != nil {
		log.Println(err)
		return Geo, err
	}
	Geo.CITY = record.City.Names["en"]
	if len(record.Subdivisions) > 0 {
		Geo.SUBDEVISIONNAME = record.Subdivisions[0].Names["en"]
	}

	Geo.IP_ADDRESS = ipaddr
	Geo.LAT = record.Location.Latitude
	Geo.LONG = record.Location.Longitude
	Geo.RADIUS = record.Location.AccuracyRadius

	return Geo, nil

}

// some of the VariableDistance content I got from gist.github.com/cdipaolo and some I wrote myself
// The part I wrote myself was the different types of measurement conversions
// I did this so I could test the method against google for accuracy.

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func VariableDistance(lat1, lon1, lat2, lon2 float64, conversiontype string) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	g := 2 * r * math.Asin(math.Sqrt(h))

	// Convert to any unit of measure needed

	if conversiontype == "meters" {
		return g
	}

	if conversiontype == "miles" {
		c := g / 1609
		return c
	}

	if conversiontype == "kilometers" {
		c := g / 1000
		return c
	}

	return g
}

func MPH(distance float64, startTime, endTime int64) int {

	st := time.Unix(startTime, 0)
	et := time.Unix(endTime, 0)
	hrs := et.Sub(st).Hours()
	mph := distance / math.Abs(hrs)

	return int(mph)
}

// This is how i found the answer to the bug issue for proceeding IP access return value in the example.
// The value in proceedingIPAcess object was the same as the POST to the endpoint.
// StartTime = (Distance / Speed) - Endtime

func GetEndTimeFromDistanceAndSpeed(distance float64, MPH float64, endTime int64) (float64, time.Time) {
	travelTime := distance / MPH
	return travelTime, time.Unix(endTime, 0).Add(-time.Hour * time.Duration(travelTime))
}

func GetStartTimeFromDistanceAndSpeed(distance float64, MPH float64, startTime int64) (float64, time.Time) {
	travelTime := distance / MPH
	return travelTime, time.Unix(startTime, 0).Add(time.Hour * time.Duration(travelTime))
}
