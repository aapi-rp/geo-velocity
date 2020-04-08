package main

import (
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

	mph := utils.MPH(vdist, 1514851200, 1514764800)

	log.Println("Miles Per Hour: ", mph)

}
