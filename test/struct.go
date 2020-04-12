package test

type DBData []struct {
	EventUUID     string `json:"event_uuid"`
	IPAddress     string `json:"ip_address"`
	UnixTimestamp int64  `json:"unix_timestamp"`
	Username      string `json:"username"`
}
