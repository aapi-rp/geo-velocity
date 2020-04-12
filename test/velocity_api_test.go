package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/aapi-rp/geo-velocity/messages"
	"github.com/aapi-rp/geo-velocity/utils"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestSuccessResponse(t *testing.T) {

	t.Log("Testing Success Status")
	status := geoAPICall(Ok200, "status", true, true, true)

	if !strings.Contains(status, "200") {
		t.Errorf("Expected a 200 response but got %s", status)
	}

	t.Log("Testing UniqueID 400 Error")
	if !strings.Contains(geoAPICall(Error400sameUUID, "body", true, true, false), messages.Err400MessageUUID) {
		t.Errorf("Expected response body message to contain %s and it did not", messages.Err400MessageUUID)
	}

	t.Log("Testing same timestamp 400 Error")
	if !strings.Contains(geoAPICall(Error400sameTimestampByUser, "body", false, true, true), messages.Err400MessageTime) {
		t.Errorf("Expected response body message to contain %s and it did not", messages.Err400MessageTime)
	}

}

func geoAPICall(jsone string, responseType string, staticUUID bool, staticUserName bool, staticUnix bool) string {
	testDBdata := DBData{}
	finalBody := ""
	json.Unmarshal([]byte(jsone), &testDBdata)
	statUser := fmt.Sprintf("%s %s", randomdata.SillyName(), randomdata.SillyName())
	uuidgenstat, erruuid := uuid.NewUUID()
	if erruuid != nil {
		logger.Error("Coud not generate UUID for testing, something went wrong: ", erruuid)
	}

	statTimeInt := utils.GetCurrentEpochTime()
	furtherRandomizeTime := statTimeInt + int64(utils.RandomNum(-40000, 40000))
	statTimeInt64 := int64(furtherRandomizeTime)

	for _, v := range testDBdata {

		ranUUID, err := uuid.NewUUID()
		randUser := fmt.Sprintf("%s %s", randomdata.SillyName(), randomdata.SillyName())

		rTimeInt := utils.GetCurrentEpochTime()
		furtherrTime := rTimeInt + int64(utils.RandomNum(-40000, 40000))
		rTimeInt64 := int64(furtherrTime)

		if staticUserName {
			v.Username = statUser
		} else {
			v.Username = randUser
		}
		if staticUUID {
			v.EventUUID = uuidgenstat.String()
		} else {
			v.EventUUID = ranUUID.String()
		}
		if staticUnix {
			v.UnixTimestamp = statTimeInt64
		} else {
			v.UnixTimestamp = rTimeInt64
		}

		endpoint := "/v1/geovelocity"
		url := utils.GetAPIUrl() + endpoint
		JsonStringData, _ := json.Marshal(v)

		var jsonStr = JsonStringData
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("There was an issue trying to call the client: ", err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if responseType == "body" {
			finalBody = string(body)
		}

		if responseType == "status" {
			finalBody = strconv.Itoa(resp.StatusCode)
		}

	}

	return finalBody
}

const Error400sameUUID = `
[
  {
    "username": "rp111",
    "unix_timestamp": 1586352357,
    "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e75",
    "ip_address": "65.60.175.99"
  },
  {
    "username": "rp111",
    "unix_timestamp": 1586438757,
    "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e75",
    "ip_address": "35.208.83.97"
  }
]
`

const Error400sameTimestampByUser = `
[
  {
    "username": "rp111",
    "unix_timestamp": 1586352357,
    "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e75",
    "ip_address": "65.60.175.99"
  },
  {
    "username": "rp111",
    "unix_timestamp": 1586438757,
    "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e76",
    "ip_address": "35.208.83.97"
  }
]
`

const Ok200 = `
[
  {
    "username": "rp111",
    "unix_timestamp": 1586352357,
    "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e75",
    "ip_address": "65.60.175.99"
  }
]
`
