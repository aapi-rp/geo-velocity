package utils

import (
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

type Coord struct {
	Latitude  float64
	Longitude float64
}

type GeoData struct {
	UUID            []byte
	LOGIN_TIME      int64
	USERNAME        string
	IP_ADDRESS      string
	LAT             float64
	LONG            float64
	RADIUS          uint16
	ISOCountryCode  string
	TIMEZONE        string
	SUBDEVISIONNAME string
	CITY            string
}

func GetGeoDataFromIP(ipaddr string) (GeoData, error) {

	Geo := GeoData{}

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

// Some of the below methods I got from gist.github.com/cdipaolo and some I wrote myself
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

	if conversiontype == "meters" {
		log.Println("meters")
		return g
	}

	if conversiontype == "miles" {
		log.Println("miles")
		c := g / 1609
		return c
	}

	if conversiontype == "kilometers" {
		log.Println("kilometers")
		c := g / 1000
		return c
	}

	return g
}

func MPH(distance float64, startTime, endTime int64) int {

	st := time.Unix(startTime, 0)
	et := time.Unix(endTime, 0)
	// Get difference in times in hours from the time stamps above
	hrs := et.Sub(st).Hours()
	// Divide distance by absolute value of hours
	sp := distance / math.Abs(hrs)

	return int(sp)
}
