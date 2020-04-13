package models

import (
	"github.com/aapi-rp/geo-velocity/db"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/model_struct"
	"github.com/aapi-rp/geo-velocity/security"
	"github.com/aapi-rp/geo-velocity/utils"
)

/*
    Method Description: Data request to add a geo event to the database.
	Parameter: geo
    Parameter Description: incoming api request data gets added to this struct
    Returns any error that may come from the database
*/
func AddEvents(geo model_struct.GeoData) error {
	err := db.InsertDBRow(db.InsertGVTable, geo.UUID, geo.LOGIN_TIME, geo.USERNAME, geo.IP_ADDRESS, geo.LAT, geo.LONG, geo.RADIUS)
	if err != nil {
		return err
	}

	return nil
}

/*
    Method Description: Validates if a specific event exists in the database that has the same unique id
	Parameter: UUID
    Parameter Description: unique id of the request
    Returns true or false if the UUID exists, and returns an error if the database has an issue
*/
func EventExistsSameUUID(UUID []byte) (bool, error) {
	uuidExists, err := db.SelectDBRowExists(db.UUIDExist, UUID)
	return uuidExists, err
}

/*
    Method Description: Validates if a specific username and timestamp combo already exists
	Parameter: user, time
    Parameter Description: user, the username of the request
    Parameter Description: time, the unix timestamp of the request
    Returns true or false if the user/time combo exists, and returns an error if the database has an issue
*/
func EventUserTimeComboExists(user string, time int64) (bool, error) {
	userTimeExists, err := db.SelectDBRowExists(db.UserLoginTimeExists, time, user)
	return userTimeExists, err
}

/*
    Method Description: Validates if a specific username and timestamp combo already exists
	Parameter: user, time
    Parameter Description: user, the username of the request
    Parameter Description: time, the unix timestamp of the request
    This method returns the final response json back to the requester
*/
func GetPreviousSubsequentCompareJSON(current model_struct.GeoData) model_struct.VelocityJSON {

	velocity := model_struct.VelocityJSON{}

	prevGeoData, hasPrevious, err := db.SelectDBRows(db.GetPrevious, current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME, current.USERNAME)

	if err != nil {
		logger.Error("Previous rows query has an issue for: ", current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)
	}

	if hasPrevious {
		vDist := utils.VariableDistance(prevGeoData.LAT, prevGeoData.LONG, current.LAT, current.LONG, "miles")
		mph := utils.MPH(vDist, prevGeoData.LOGIN_TIME, current.LOGIN_TIME)
		decryptedIP, _ := security.Decrypt(prevGeoData.IP_ADDRESS)
		velocity.PrecedingIPAccess.IP = decryptedIP
		velocity.PrecedingIPAccess.Lat = prevGeoData.LAT
		velocity.PrecedingIPAccess.Lon = prevGeoData.LONG
		velocity.PrecedingIPAccess.Radius = prevGeoData.RADIUS
		velocity.PrecedingIPAccess.Timestamp = prevGeoData.LOGIN_TIME
		velocity.PrecedingIPAccess.Speed = int64(mph)

		if mph > 500 {
			velocity.TravelToCurrentGeoSuspicious = true
		} else {
			velocity.TravelToCurrentGeoSuspicious = false
		}
	}

	subGeoData, hasSubsequent, err := db.SelectDBRows(db.GetSubsequent, current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME, current.USERNAME)

	if err != nil {
		logger.Error("Subsequent rows query has an issue for: ", current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)
	}

	if hasSubsequent {
		vDist := utils.VariableDistance(subGeoData.LAT, subGeoData.LONG, current.LAT, current.LONG, "miles")
		mph := utils.MPH(vDist, current.LOGIN_TIME, subGeoData.LOGIN_TIME)
		decryptedIP, _ := security.Decrypt(subGeoData.IP_ADDRESS)
		velocity.SubsequentIPAccess.IP = decryptedIP
		velocity.SubsequentIPAccess.Lat = subGeoData.LAT
		velocity.SubsequentIPAccess.Lon = subGeoData.LONG
		velocity.SubsequentIPAccess.Radius = subGeoData.RADIUS
		velocity.SubsequentIPAccess.Timestamp = subGeoData.LOGIN_TIME
		velocity.SubsequentIPAccess.Speed = int64(mph)

		if mph > 500 {
			velocity.TravelFromCurrentGeoSuspicious = true
		} else {
			velocity.TravelFromCurrentGeoSuspicious = false
		}
	}

	velocity.CurrentGeo.Radius = current.RADIUS
	velocity.CurrentGeo.Lat = current.LAT
	velocity.CurrentGeo.Lon = current.LONG

	return velocity

}
