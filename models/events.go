package models

import (
	"github.com/aapi-rp/geo-velocity/db"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/model_struct"
	"github.com/aapi-rp/geo-velocity/security"
	"github.com/aapi-rp/geo-velocity/utils"
)

func AddEvents(geo model_struct.GeoData) error {
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

func GetPreviousSubsequentCompareJSON(current model_struct.GeoData) model_struct.VelocityJSON {

	velocity := model_struct.VelocityJSON{}

	prevGeoData, hasPrevious, err := db.SelectDBRows(db.GetPrevious, current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)

	if err != nil {
		logger.Error("Previous rows query has an issue for: ", current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)
	}

	if hasPrevious {
		vdist := utils.VariableDistance(prevGeoData.LAT, prevGeoData.LONG, current.LAT, current.LONG, "miles")
		mph := utils.MPH(vdist, prevGeoData.LOGIN_TIME, current.LOGIN_TIME)
		decryptedIP, _ := security.Decrypt(prevGeoData.IP_ADDRESS)
		velocity.PrecedingIPAccess.IP = decryptedIP
		velocity.PrecedingIPAccess.Lat = prevGeoData.LAT
		velocity.PrecedingIPAccess.Lon = prevGeoData.LONG
		velocity.PrecedingIPAccess.Radius = prevGeoData.RADIUS
		velocity.PrecedingIPAccess.Speed = int64(mph)

		if mph > 500 {
			velocity.TravelToCurrentGeoSuspicious = true
		} else {
			velocity.TravelToCurrentGeoSuspicious = false
		}
	}

	subGeoData, hasSubsequent, err := db.SelectDBRows(db.GetSubsequent, current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)

	if err != nil {
		logger.Error("Subsequent rows query has an issue for: ", current.IP_ADDRESS, current.USERNAME, current.LOGIN_TIME)
	}

	if hasSubsequent {
		vdist := utils.VariableDistance(subGeoData.LAT, subGeoData.LONG, current.LAT, current.LONG, "miles")
		mph := utils.MPH(vdist, current.LOGIN_TIME, subGeoData.LOGIN_TIME)
		decryptedIP, _ := security.Decrypt(subGeoData.IP_ADDRESS)
		velocity.SubsequentIPAccess.IP = decryptedIP
		velocity.SubsequentIPAccess.Lat = subGeoData.LAT
		velocity.SubsequentIPAccess.Lon = subGeoData.LONG
		velocity.SubsequentIPAccess.Radius = subGeoData.RADIUS
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
