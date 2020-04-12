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

/*
   Method Description: This function will test the API for specific success and error responses.
*/

func TestSuccessErrorResponses(t *testing.T) {

	t.Log("Testing Success Status")
	status := geoAPICallHarness(Ok200, "status", true, true, true, false, "")

	if !strings.Contains(status, "200") {
		t.Errorf("Expected a 200 response but got %s", status)
	}

	t.Log("Testing UniqueID 400 Error")
	if !strings.Contains(geoAPICallHarness(Error400sameUUID, "body", true, true, false, false, ""), messages.Err400MessageUUID) {
		t.Errorf("Expected response body message to contain %s and it did not", messages.Err400MessageUUID)
	}

	t.Log("Testing same timestamp 400 Error")
	if !strings.Contains(geoAPICallHarness(Error400sameTimestampByUser, "body", false, true, true, false, ""), messages.Err400MessageTime) {
		t.Errorf("Expected response body message to contain %s and it did not", messages.Err400MessageTime)
	}
}

/*
   Method Description: This function will dynamically test the example in the Challenge
*/
func TestResponseDataOldSchoolSecureWorks(t *testing.T) {

	statUser := fmt.Sprintf("%s %s", randomdata.SillyName(), randomdata.SillyName())
	subUser := fmt.Sprintf("%s %s", randomdata.SillyName(), randomdata.SillyName())

	t.Log("Testing geo data first entry no data")
	body1 := geoAPICallHarness(oldSchoolSecureWorks1, "body", false, true, true, true, statUser)

	if !strings.Contains(body1, oldSchoolResponse1) {
		t.Errorf("Expected body of %s, but got %s", oldSchoolResponse1, body1)
	}

	t.Log("Testing geo data second entry with only preceding data")
	body2 := geoAPICallHarness(oldSchoolSecureWorks2, "body", false, true, true, true, statUser)

	if !strings.Contains(body2, oldSchoolResponse2) {
		t.Errorf("Expected body of %s, but got %s", oldSchoolResponse2, body2)
	}

	t.Log("Testing geo data third entry with both preceding and subsequent data")
	body3 := geoAPICallHarness(oldSchoolSecureWorks3, "body", false, true, true, true, statUser)

	if !strings.Contains(body3, oldSchoolResponse3) {
		t.Errorf("Expected body of %s, but got %s", oldSchoolResponse3, body3)
	}

	t.Log("Challenge old school response is: ", body3)

	t.Log("Testing geo data forth entry with no data")
	_ = geoAPICallHarness(subSequent1, "body", false, true, true, true, subUser)

	t.Log("Testing geo data fith entry with subsequent data and suspicious activity")
	body5 := geoAPICallHarness(subSequent2, "body", false, true, true, true, subUser)

	t.Log("Subsequent result with Suspicious activity: ", body5)

	if !strings.Contains(body5, subSequentResponseWithSuspiciousActivity) {
		t.Errorf("Expected body of %s, but got %s", subSequentResponseWithSuspiciousActivity, body5)
	}
}

/*
   Method Description: This testing harness allows for dynamic data on top of json, so you can send it json
   and overwrite specific values of the request with dynamic data, or keep some
   dynamic and some static its up to you how you want to use it, if I had more time
   I would do some fun stuff this.
*/

func geoAPICallHarness(jsone string, responseType string, staticUUID bool, staticUserName bool, staticUnix bool, usejsontimestamp bool, hrdCodedName string) string {
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

		if hrdCodedName != "" {
			v.Username = hrdCodedName
		}

		if staticUserName {
			if hrdCodedName == "" {
				v.Username = statUser
			}
		} else {
			if hrdCodedName == "" {
				v.Username = randUser
			}
		}
		if staticUUID {
			v.EventUUID = uuidgenstat.String()
		} else {
			v.EventUUID = ranUUID.String()
		}
		if staticUnix {
			if !usejsontimestamp {
				v.UnixTimestamp = statTimeInt64
			}
		} else {
			if !usejsontimestamp {
				v.UnixTimestamp = rTimeInt64
			}
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

const oldSchoolSecureWorks1 = `
[
  	{
    	"username": "bob",
    	"unix_timestamp": 1514678400,
    	"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e75",
    	"ip_address": "24.242.71.20"
  	}
]
`

const oldSchoolSecureWorks2 = `
[
  	{
    	"username": "bob",
    	"unix_timestamp": 1514851200,
    	"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e76",
    	"ip_address": "91.207.175.104"
  	}
]
`

const oldSchoolSecureWorks3 = `
[
	{
    	"username": "bob",
    	"unix_timestamp": 1514764800,
    	"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e77",
    	"ip_address": "206.81.252.6"
	}
]
`

const subSequent1 = `
[
	{
    	"username": "bob",
    	"unix_timestamp": 1586658993,
    	"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e77",
    	"ip_address": "70.162.226.163"
	}
]
`

const subSequent2 = `
[
	{
    	"username": "bob",
    	"unix_timestamp": 1586657121,
    	"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e77",
    	"ip_address": "107.175.170.77"
	}
]
`

const oldSchoolResponse1 = `{"currentGeo":{"lat":30.3773,"lon":-97.71,"radius":5},"precedingIpAccess":{},"subsequentIpAccess":{},"travelFromCurrentGeoSuspicious":false,"travelToCurrentGeoSuspicious":false}`
const oldSchoolResponse2 = `{"currentGeo":{"lat":34.0549,"lon":-118.2578,"radius":200},"precedingIpAccess":{"ip":"24.242.71.20","lat":30.3773,"lon":-97.71,"radius":5,"speed":25,"timestamp":1514678400},"subsequentIpAccess":{},"travelFromCurrentGeoSuspicious":false,"travelToCurrentGeoSuspicious":false}`
const oldSchoolResponse3 = `{"currentGeo":{"lat":38.9206,"lon":-76.8787,"radius":50},"precedingIpAccess":{"ip":"24.242.71.20","lat":30.3773,"lon":-97.71,"radius":5,"speed":55,"timestamp":1514678400},"subsequentIpAccess":{"ip":"91.207.175.104","lat":34.0549,"lon":-118.2578,"radius":200,"speed":96,"timestamp":1514851200},"travelFromCurrentGeoSuspicious":false,"travelToCurrentGeoSuspicious":false}`
const subSequentResponseWithSuspiciousActivity = `{"currentGeo":{"lat":37.751,"lon":-97.822,"radius":1000},"precedingIpAccess":{},"subsequentIpAccess":{"ip":"70.162.226.163","lat":33.491,"lon":-112.2491,"radius":10,"speed":1657,"timestamp":1586658993},"travelFromCurrentGeoSuspicious":true,"travelToCurrentGeoSuspicious":false}`
