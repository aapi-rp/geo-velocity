package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func main() {
	db, err := geoip2.Open("db/GeoLite2-City.mmdb")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	//67.181.148.192
	ip := net.ParseIP("72.208.39.180")
	record, err := db.City(ip)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("City name: %v\n", record.City.Names["en"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}

	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

}
