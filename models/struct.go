package models

type Coord struct {
	Latitude  float64
	Longitude float64
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
