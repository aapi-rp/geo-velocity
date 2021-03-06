package messages

import (
	"encoding/json"
	"github.com/aapi-rp/geo-velocity/model_struct"
)

const Err500Message = "Error, Something went wrong with your request, please contact your administrator."
const Err400Message = "Error, Malformed Request"
const Err500MessageContentTypeJson = "Error, your request should be adjusted to add Content-Type header with mime application/json"
const Err400MessageUUID = "Error, UUID Exists"
const Err400MessageTime = "Error, User event already exists with that timestamp"

/*
    Method Description: Creates a json message form text sent.
	Parameter: message
    Parameter Description: the message of the json
    Parameter: status
    Parameter Description: the http status of the message
    Returns any error that may come from the database
*/
func CreateJsonMssage(message string, status string) ([]byte, error) {
	webMessage := model_struct.WebMessage{}
	webMessage.Status = status
	webMessage.Message = message
	jsonMessage, err := json.Marshal(webMessage)
	return jsonMessage, err
}
