package models

type Coord struct {
	Latitude  float64
	Longitude float64
}

type WebMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type GeoData struct {
	UUID            []byte
	LOGIN_TIME      int64
	USERNAME        string
	IP_ADDRESS      string
	LAT             float64
	LONG            float64
	RADIUS          uint16
	ISOCountryCode  string
	TIMEZONE        string
	SUBDEVISIONNAME string
	CITY            string
}

type BaseEventRequest struct {
	EventUUID     string `json:"event_uuid"`
	IPAddress     string `json:"ip_address"`
	UnixTimestamp int64  `json:"unix_timestamp"`
	Username      string `json:"username"`
}
