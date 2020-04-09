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

func EventExistsSameUUID(UUID []byte) (bool, error) {
	uuidExists, err := db.SelectDBRowExists(db.UUIDExist, UUID)
	return uuidExists, err
}

func EventUserTimeComboExists(user string, time int64) (bool, error) {
	userTimeExists, err := db.SelectDBRowExists(db.UserLoginTimeExists, time, user)
	return userTimeExists, err
}
