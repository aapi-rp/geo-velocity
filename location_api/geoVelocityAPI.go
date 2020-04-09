package location_api

import (
	"encoding/json"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/messages"
	"github.com/aapi-rp/geo-velocity/models"
	"github.com/aapi-rp/geo-velocity/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func EventData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error getting request body: ", err)
		msg, _ := messages.CreateJsonMssage(messages.Err400Message, "400")
		w.WriteHeader(500)
		w.Write(msg)
		return
	}
	log.Println(string(body))
	var er models.BaseEventRequest
	err = json.Unmarshal(body, &er)
	if err != nil {
		logger.Error("Error Unmarshaling json into struct: ", err)
		msg, _ := messages.CreateJsonMssage(messages.Err400Message, "400")
		w.WriteHeader(500)
		w.Write(msg)
		return
	}

	geo, err := utils.GetGeoDataFromIP(er.IPAddress)

	geo.USERNAME = er.Username
	geo.LOGIN_TIME = er.UnixTimestamp
	geo.IP_ADDRESS = er.IPAddress

}
