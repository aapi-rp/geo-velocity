package location_api

import (
	"encoding/json"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/messages"
	"github.com/aapi-rp/geo-velocity/models"
	"github.com/aapi-rp/geo-velocity/security"
	"github.com/aapi-rp/geo-velocity/utils"
	"io/ioutil"
	"net/http"
)

func EventData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error getting request body: ", err)
		msg, _ := messages.CreateJsonMssage(messages.Err400Message, "400")
		w.WriteHeader(400)
		w.Write(msg)
		return
	}

	var er models.BaseEventRequest
	err = json.Unmarshal(body, &er)
	if err != nil {
		logger.Error("Error getting request body: ", err)
		msg, _ := messages.CreateJsonMssage(messages.Err400Message, "400")
		w.WriteHeader(400)
		w.Write(msg)
		return
	}

	geo, err := utils.GetGeoDataFromIP(er.IPAddress)

	geo.USERNAME = er.Username
	geo.LOGIN_TIME = er.UnixTimestamp
	geo.IP_ADDRESS = security.Encrypt(er.IPAddress)
	geo.UUID = []byte(er.EventUUID)

	uuidExists, err := models.EventExistsSameUUID(geo.UUID)

	if uuidExists {
		logger.Error("Error, UUID Exists: ", err)
		msg, _ := messages.CreateJsonMssage(messages.Err400MessageUUID, "400")
		w.WriteHeader(400)
		w.Write(msg)
		return
	} else {
		userTimeComboExists, err := models.EventUserTimeComboExists(geo.USERNAME, geo.LOGIN_TIME)
		if userTimeComboExists {
			logger.Error("Error, time user combo already exists: ", err)
			msg, _ := messages.CreateJsonMssage(messages.Err400MessageTime, "400")
			w.WriteHeader(400)
			w.Write(msg)
			return
		} else {
			err = models.AddEvents(geo)
			if err != nil {
				logger.Error("DB insert event failed: ", err)
				msg, _ := messages.CreateJsonMssage(messages.Err500Message, "500")
				w.WriteHeader(500)
				w.Write(msg)
				return
			}

			// Start forming return JSON, we need previous/subsequent and current

		}
	}
}
