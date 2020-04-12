package utils

import (
	"fmt"
	"github.com/aapi-rp/geo-velocity/config"
	"github.com/aapi-rp/geo-velocity/model_struct"
	"github.com/oschwald/geoip2-golang"
	"log"
	"math"
	"math/rand"
	"net"
	"time"
)

/*
   Method Description: Creates a unix timestamp in seconds
   Returns an int64 unix timestamp
*/
func GetCurrentEpochTime() int64 {
	now := time.Now()
	secs := now.Unix()
	return secs
}

/*
    Method Description: Creates a random number between min and max
	Parameter: min
    Parameter Description: Minimum number that the random number can be
    Parameter: max
    Parameter Description: Maximum number that the random number can be
    Returns a random int between min and max parameters
*/
func RandomNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/*
   Method Description: Concatenate environment variables to generate this servers api url
   Returns this servers api url
*/
func GetAPIUrl() string {
	URL := ""
	if config.ServerPort() == "80" || config.ServerPort() == "443" {
		URL = fmt.Sprintf("%s://%s", config.ServerScheme(), config.ServerHost())
	} else {
		URL = fmt.Sprintf("%s://%s:%s", config.ServerScheme(), config.ServerHost(), config.ServerPort())
	}
	return URL
}

/*
    Method Description: Get longitude and latitude from IP address
	Parameter: ipaddr
    Parameter Description: The IP address you would like to get long and lat from
    Returns lat and long data
*/
func GetGeoDataFromIP(ipaddr string) (model_struct.GeoData, error) {

	Geo := model_struct.GeoData{}

	db, err := geoip2.Open("data/GeoLite2-City.mmdb")
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

/*

    Some of the VariableDistance content I got from gist.github.com/cdipaolo and some I wrote myself
    The part I wrote myself was the different types of measurement conversions
    I did this so I could test the method against google for accuracy.

    Method Description: A component of the Haversine formula
	Parameter: theta
    Parameter Description: Part of the calculation
    Returns math.Pow

    In mathematics, the sine is a trigonometric function of an angle.
    The sine of an acute angle is defined in the context of a right triangle:
    for the specified angle, it is the ratio of the length of the side that is
    opposite that angle to the length of the longest side of the triangle (the hypotenuse)
*/
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

/*

    Some of the VariableDistance content I got from gist.github.com/cdipaolo and some I wrote myself
    The part I wrote myself was the different types of measurement conversions
    I did this so I could test the method against google for accuracy.

    Method Description: Generates a float of the distance between 2 geographic locations
	Parameter: lat1, lon1
    Parameter Description: The first coordinance of the geo location
    Parameter: lat2, lon2
    Parameter Description: The second coordinance of the geo location
    Parameter: conversionType
    Parameter Description: The measurement type you want back from the method meters, miles, or kilometers
    Returns distance in meters, miles, or kilometers

    The haversine formula determines the great-circle distance between two points on a sphere given their longitudes and latitudes.
    Important in navigation, it is a special case of a more general formula in spherical trigonometry, the law of haversines,
    that relates the sides and angles of spherical triangles.
*/
func VariableDistance(lat1, lon1, lat2, lon2 float64, conversionType string) float64 {
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

	if conversionType == "meters" {
		return g
	}

	if conversionType == "miles" {
		c := g / 1609
		return c
	}

	if conversionType == "kilometers" {
		c := g / 1000
		return c
	}

	return g
}

/*
    Method Description: Get Miles Per Hour from distance start time and end time
	Parameter: distance
    Parameter Description: The distance between 2 points
	Parameter: startTime
    Parameter Description: unix timestamp of the start time
	Parameter: endTime
    Parameter Description: unix timestamp of the end time
    Returns Miles Per Hour based on the the parameters provided
*/
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

/*
    Method Description: Calculate unix timestamp for start time based on distance, speed, and end time
	Parameter: distance
    Parameter Description: The distance between 2 points
	Parameter: MPH
    Parameter Description: the Miles Per Hour
	Parameter: endTime
    Parameter Description: unix timestamp of the end time
    Returns End Time in Unix Seconds
*/

func GetEndTimeFromDistanceAndSpeed(distance float64, MPH float64, startTime int64) (float64, time.Time) {
	travelTime := distance / MPH
	return travelTime, time.Unix(startTime, 0).Add(-time.Hour * time.Duration(travelTime))
}

/*
    Method Description: Calculate unix timestamp for end time based on distance, speed, and start time
	Parameter: distance
    Parameter Description: The distance between 2 points
	Parameter: MPH
    Parameter Description: the Miles Per Hour
	Parameter: startTime
    Parameter Description: unix timestamp of the start time
    Returns Start Time in Unix Seconds
*/

func GetStartTimeFromDistanceAndSpeed(distance float64, MPH float64, endTime int64) (float64, time.Time) {
	travelTime := distance / MPH
	return travelTime, time.Unix(endTime, 0).Add(time.Hour * time.Duration(travelTime))
}
