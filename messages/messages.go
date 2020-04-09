package messages

import (
	"encoding/json"
	"github.com/aapi-rp/geo-velocity/models"
)

const Err500Message = "Error, Something went wrong with your request, please contact your administrator."
const Err400Message = "Error, Malformed Request"
const Err400MessageUUID = "Error, UUID Exists"
const Err400MessageTime = "Error, User event already exists with that timestamp"

func CreateJsonMssage(message string, status string) ([]byte, error) {
	webmessage := models.WebMessage{}
	webmessage.Status = status
	webmessage.Message = message
	jsonMessage, err := json.Marshal(webmessage)
	return jsonMessage, err
}
