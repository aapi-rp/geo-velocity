# Geo Velocity

##### Detailed Documentation
- [Geo Velocity Home](https://github.com/aapi-rp/geo-velocity/wiki "Geo Velocity Home")
- [Requirements Abstract](https://github.com/aapi-rp/geo-velocity/wiki/Requirements_Abstract "Requirements Abstract")
- [Geo API Document](https://github.com/aapi-rp/geo-velocity/wiki/API-Docs)

What is geo Velocity?

Geo velocity is The process in which an IP address is captured by standard programmatic means, and that IP address is analyzed for its relevant geo location, more specifically, the origin of the IP addresses longitude and latitude. When the IP address is captured, the time in which it was captured is also collected in some type of storage table for later comparison. If you have 2 different geo access points (2 different geographic locations) captured in this data storage table, you can then compare the times of access, and the distance between the access points to get the Miles Per Hour or Kilometers Per Hour it would take to get from one of the geographic locations to the other. With the speed of travel defined, you can tell if the travel between the locations would be feasible, if its not feasible, then you can alert someone via email, or simply programmatically block the access.

What is the purpose of this project?

To provide an API that calculates geo events based on IP Address origin, and can decipher if the traffic is suspicious or valid, and return deciphered results to any application that needs to protect against malicious attacks based on geographic location, time and speed.

## Docker
#### Pull docker from dockerhub
```
$ docker pull aapirp/geo-velocity:v1.1.5
$ docker run --publish 8081:8081 -e server_port=8081 --name geo aapirp/geo-velocity:v1.1.5
```

#### Environment Variables

All settings below during testing are defaulted and do not need to be changed unless using in production mode. Keys should be added as kubernetes secrets in the GCP interface for security. [Kubernetes Secrets](https://kubernetes.io/docs/concepts/configuration/secret/ "Kubernetes Secrets")

Terminal Command to add env variables:
```
$ export sqlite3_db_path=data/geo-velocity.sqlite3
$ export env=development | production
$ export encryption_key=256 hex key
$ export encryption_iv=256 hex iv
$ export enable_ssl=true
$ export server_port=8080
$ export skip_ssl_verify
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
$ cd to /yourbase/github.com/aapi-rp
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

#### Request Parameters

| Name           | Example                              | Type | Required | Data Type |
|----------------|--------------------------------------|------|----------|-----------|
| username       | jim                                  | Form | yes      | string    |
| unix_timestamp | 1586477927                           | Form | yes      | number    |
| event_uuid     | 85ad929a-db03-4bf4-9541-8f728fa12e98 | Form | yes      | string    |
| ip_address     | 65.49.22.66                          | Form | yes      | striing   |


#### Example Request:

```
$ curl -X POST -d '{"username": "jim",
"unix_timestamp": 1586477927,
"event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e98", "ip_address": "65.49.22.66"}' http://localhost:8080/v1/geovelocity
```

#### Example Response:

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
