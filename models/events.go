package models

import (
	"github.com/aapi-rp/geo-velocity/db"
	"github.com/aapi-rp/geo-velocity/logger"
)

func AddEventsTestingFixtures(geo GeoData) {

	err := db.InsertDBRow(db.InsertGVTable, geo.UUID, geo.LOGIN_TIME, geo.USERNAME, geo.IP_ADDRESS, geo.LAT, geo.LONG, geo.RADIUS)

	if err != nil {
		logger.Error("Prepare insert event failed: ", err)
	}
}
