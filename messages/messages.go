package messages

import (
	"encoding/json"
	"github.com/aapi-rp/geo-velocity/models"
)

const Err500Message = "Something went wrong with your request, please contact your administrator."
const Err400Message = "Malformed Request"

func CreateJsonMssage(message string, status string) ([]byte, error) {
	webmessage := models.WebMessage{}
	webmessage.Status = status
	webmessage.Message = message
	jsonMessage, err := json.Marshal(webmessage)
	return jsonMessage, err
}
