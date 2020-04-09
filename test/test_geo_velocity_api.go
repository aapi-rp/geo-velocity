package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postRequestGeoVelocityAPI() {
	url := "http://localhost:8080/v1/geovelocity"
	fmt.Println("URL: ", url)

	var jsonStr = []byte(testjson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

const testjson = `
{
  "username": "bob",
  "unix_timestamp": 1514764800,
  "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e42",
  "ip_address": "206.81.252.6"
}`
