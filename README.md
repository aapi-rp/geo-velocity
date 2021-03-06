# Geo Velocity

##### Detailed Documentation
- [Geo Velocity Home](https://github.com/aapi-rp/geo-velocity/wiki "Geo Velocity Home")
- [Requirements Abstract](https://github.com/aapi-rp/geo-velocity/wiki/Requirements_Abstract "Requirements Abstract")
- [Geo API Document](https://github.com/aapi-rp/geo-velocity/wiki/API-Docs)
- [Acknowledgments](https://github.com/aapi-rp/geo-velocity/wiki/Acknowledgments)

 
What is geo Velocity?

Geo velocity is The process in which an IP address is captured by standard programmatic means, and that IP address is analyzed for its relevant geo location, more specifically, the origin of the IP addresses longitude and latitude. When the IP address is captured, the time in which it was captured is also collected in some type of storage table for later comparison. If you have 2 different geo access points (2 different geographic locations) captured in this data storage table, you can then compare the times of access, and the distance between the access points to get the Miles Per Hour or Kilometers Per Hour it would take to get from one of the geographic locations to the other. With the speed of travel defined, you can tell if the travel between the locations would be feasible, if its not feasible, then you can alert someone via email, or simply programmatically block the access.

What is the purpose of this project?

To provide an API that calculates geo events based on IP Address origin, and can decipher if the traffic is suspicious or valid, and return deciphered results to any application that needs to protect against malicious attacks based on geographic location, time and speed.

## Prerequisites

* Docker 19.03.8, build afacb8b or later, I recommend [Docker Desktop](https://www.docker.com/products/docker-desktop) if its not installed
* Golang 1.13.8 - If manual install
* An understanding of API's

## Docker
#### Pull docker from dockerhub
```
$ docker pull aapirp/geo-velocity:v1.1.8
$ docker run --publish 8081:8081 -e server_port=8081 --name geo aapirp/geo-velocity:v1.1.8
```

#### Environment Variables

All settings below during testing are defaulted and do not need to be changed unless using in production mode, or if you are already using port 8081. Keys should be added as kubernetes secrets in the GCP interface for security. [Kubernetes Secrets](https://kubernetes.io/docs/concepts/configuration/secret/ "Kubernetes Secrets")

Terminal Command to add env variables:
```
$ export sqlite3_db_path=data/geo-velocity.sqlite3
$ export env=development | production
$ export encryption_key=256 hex key
$ export encryption_iv=256 hex iv
$ export enable_ssl=true
$ export server_port=8080
$ export skip_ssl_verify=false
$ export server_host=localhost
$ export server_scheme=https
```

#### Build your own docker

run the following:
```
$ git clone https://github.com/aapi-rp/geo-velocity.git
$ cd to /yourbase/github.com/aapi-rp
$ docker build geo-velocity
```

## Install from source
run the following:
```
$ git clone https://github.com/aapi-rp/geo-velocity.git
$ cd to /yourbase/github.com/aapi-rp/geo-velocity
$ go run main.go
```

## Testing

```
$ cd test
$ go test -v
```

Example testing output:

```
=== RUN   TestSuccessErrorResponses
--- PASS: TestSuccessErrorResponses (0.01s)
    velocity_api_test.go:25: Testing Success Status
    velocity_api_test.go:32: Testing UniqueID 400 Error
    velocity_api_test.go:37: Testing same timestamp 400 Error
=== RUN   TestResponseDataOldSchoolSecureWorks
--- PASS: TestResponseDataOldSchoolSecureWorks (0.01s)
    velocity_api_test.go:51: Testing geo data first entry no data
    velocity_api_test.go:58: Testing geo data second entry with only preceding data
    velocity_api_test.go:65: Testing geo data third entry with both preceding and subsequent data
    velocity_api_test.go:72: Challenge old school response is:  {"currentGeo":{"lat":38.9206,"lon":-76.8787,"radius":50},"precedingIpAccess":{"ip":"24.242.71.20","lat":30.3773,"lon":-97.71,"radius":5,"speed":55,"timestamp":1514678400},"subsequentIpAccess":{"ip":"91.207.175.104","lat":34.0549,"lon":-118.2578,"radius":200,"speed":96,"timestamp":1514851200},"travelFromCurrentGeoSuspicious":false,"travelToCurrentGeoSuspicious":false}
    velocity_api_test.go:74: Testing geo data forth entry with no data
    velocity_api_test.go:77: Testing geo data fifth entry with subsequent data and suspicious activity
    velocity_api_test.go:80: Subsequent result with Suspicious activity:  {"currentGeo":{"lat":37.751,"lon":-97.822,"radius":1000},"precedingIpAccess":{},"subsequentIpAccess":{"ip":"70.162.226.163","lat":33.491,"lon":-112.2491,"radius":10,"speed":1657,"timestamp":1586658993},"travelFromCurrentGeoSuspicious":true,"travelToCurrentGeoSuspicious":false}
PASS
ok  	github.com/aapi-rp/geo-velocity/test	0.232s
```


## Geo Velocity API

#### MIME Type
`
application/json
`
#### Request Type
`
POST
`

#### Request JSON Parameters

| Name           | Example                              | Required | Data Type |
|----------------|--------------------------------------|----------|-----------|
| username       | jim                                  | yes      | string    |
| unix_timestamp | 1586477927                           | yes      | number    |
| event_uuid     | 85ad929a-db03-4bf4-9541-8f728fa12e98 | yes      | string    |
| ip_address     | 65.49.22.66                          | yes      | striing   |


#### Example Request:

```
$ curl --location --request POST 'localhost:8081/v1/geovelocity' --header 'Content-Type: application/json' --data-raw '{ "username": "jim", "unix_timestamp": 1586477934, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e934","ip_address": "65.49.22.66"}'
```

#### Example Response:


Status: <span style="color:green">200 ok</span>

```
{
  "currentGeo": {
    "lat": 33.491,
    "lon": -112.2491,
    "radius": 10
  },
  "precedingIpAccess": {
    "ip": "35.208.83.97",
    "lat": 37.751,
    "lon": -97.822,
    "radius": 1000,
    "speed": 79,
    "timestamp": 1586438757
  },
  "subsequentIpAccess": {
    "ip": "78.31.205.251",
    "lat": 40.7308,
    "lon": -73.9975,
    "radius": 1000,
    "speed": 163,
    "timestamp": 1586525157
  },
  "travelFromCurrentGeoSuspicious": false,
  "travelToCurrentGeoSuspicious": false
}
```

Status: <span style="color:red">400 Bad Request</span>


This error is farily clear, its when a UniqueID exists and someone trys to use it again.
```
{
  "message": "Error, UUID Exists",
  "status": "400"
}
```
This happens because a specific user has the submitted unix timestamp already.
```
{
  "message": "Error, User event already exists with that timestamp",
  "status": "400"
}
```
This happens when you are sending bad json.
```
{
  "message": "Error, Malformed Request",
  "status": "400"
}
```

Status: <span style="color:red">500 Internal Server Error</span>

This happens because you are missing the mime type application/json in your Content-Type header.
```
{
  "message": "Error, your request should be adjusted to add Content-Type header with mime application/json",
  "status": "500"
}
```
This is just bad stuff happening on our side that one of our devs messed up.  Forgive them they are great people... seriously.
```
{
  "message": "Error, Something went wrong with your request, please contact your administrator.",
  "status": "500"
}
```

