package main

import (
	"github.com/aapi-rp/geo-velocity/db"
	"github.com/aapi-rp/geo-velocity/utils"
	"log"
)

func main() {
	geoP, err := utils.GetGeoDataFromIP("91.207.175.104")

	if err != nil {
		log.Println(err)
	}

	log.Println(geoP.LAT, geoP.LONG, geoP.CITY, geoP.SUBDEVISIONNAME)

	curr := utils.Coord{}

	curr.Longitude = geoP.LONG
	curr.Latitude = geoP.LAT

	geoC, err := utils.GetGeoDataFromIP("206.81.252.6")

	if err != nil {
		log.Println(err)
	}

	log.Println(geoC.LAT, geoC.LONG, geoC.CITY, geoC.SUBDEVISIONNAME)

	pre := utils.Coord{}

	pre.Longitude = geoC.LONG
	pre.Latitude = geoC.LAT

	vdist := utils.VariableDistance(geoP.LAT, geoP.LONG, geoC.LAT, geoC.LONG, "miles")

	log.Println("Variable Distance: ", vdist)

	mph := utils.MPH(vdist, 1514764800, 1514851200)

	endhours, timestamp := utils.GetEndTimeFromDistanceAndSpeed(vdist, float64(mph), 1514851200)

	starthours, timestamp2 := utils.GetStartTimeFromDistanceAndSpeed(vdist, float64(mph), 1514764800)

	log.Println("Hours the trip would take: ", endhours, "hours, and ending timestamp of: ", timestamp.Unix())

	log.Println("Hours the trip would take: ", starthours, "hours and starting timestamp of: ", timestamp2.Unix())

	log.Println("Miles Per Hour: ", mph)

	db.InitData()

}

// Answer to the bug is 1514678400
