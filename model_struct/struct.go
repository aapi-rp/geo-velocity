package model_struct

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
	MPH             int
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

type VelocityJSON struct {
	CurrentGeo struct {
		Lat    float64 `json:"lat,omitempty"`
		Lon    float64 `json:"lon,omitempty"`
		Radius uint16  `json:"radius,omitempty"`
	} `json:"currentGeo,omitempty"`
	PrecedingIPAccess struct {
		IP        string  `json:"ip,omitempty"`
		Lat       float64 `json:"lat,omitempty"`
		Lon       float64 `json:"lon,omitempty"`
		Radius    uint16  `json:"radius,omitempty"`
		Speed     int64   `json:"speed,omitempty"`
		Timestamp int64   `json:"timestamp,omitempty"`
	} `json:"precedingIpAccess,omitempty"`
	SubsequentIPAccess struct {
		IP        string  `json:"ip,omitempty"`
		Lat       float64 `json:"lat,omitempty"`
		Lon       float64 `json:"lon,omitempty"`
		Radius    uint16  `json:"radius,omitempty"`
		Speed     int64   `json:"speed,omitempty"`
		Timestamp int64   `json:"timestamp,omitempty"`
	} `json:"subsequentIpAccess,omitempty"`
	TravelFromCurrentGeoSuspicious bool `json:"travelFromCurrentGeoSuspicious"`
	TravelToCurrentGeoSuspicious   bool `json:"travelToCurrentGeoSuspicious"`
}
