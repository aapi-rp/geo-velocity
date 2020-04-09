package models

import (
	"github.com/aapi-rp/geo-velocity/db"
)

func AddEvents(geo GeoData) error {
	err := db.InsertDBRow(db.InsertGVTable, geo.UUID, geo.LOGIN_TIME, geo.USERNAME, geo.IP_ADDRESS, geo.LAT, geo.LONG, geo.RADIUS)
	if err != nil {
		return err
	}

	return nil
}
