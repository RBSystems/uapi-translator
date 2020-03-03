package models

import "github.com/byuoitav/common/structs"

type CouchSearch struct {
	GT    string `json:"$gt,omitempty"`
	LT    string `json:"$lt,omitempty"`
	Regex string `json:"$regex,omitempty"`
}

type RoomQuery struct {
	Selector struct {
		ID CouchSearch `json:"_id"`
	} `json:"selector"`
	Limit int `json:"limit"`
}

type DeviceTypeQuery struct {
	ID *CouchSearch `json:"_id,omitempty"`
}

type DeviceQuery struct {
	Selector struct {
		ID      CouchSearch      `json:"_id"`
		DevType *DeviceTypeQuery `json:"type,omitempty"`
	} `json:"selector"`
	Limit int `json:"limit"`
}

type DisplayQuery struct {
	Selector struct {
		ID CouchSearch `json:"_id"`
	} `json:"selector"`
	Limit int `json:"limit"`
}

// Rooms
type RoomResponse struct {
	Docs     []RoomDB `json:"docs"`
	Bookmark string   `json:"bookmark"`
	Warning  string   `json:"warning"`
}

type RoomDB struct {
	Rev string `json:"_rev,omitempty"`
	*structs.Room
}

type DeviceResponse struct {
	Docs     []DeviceDB `json:"docs"`
	Bookmark string     `json:"bookmark"`
	Warning  string     `json:"warning"`
}

type DeviceDB struct {
	Rev string `json:"_rev,omitempty"`
	*structs.Device
}
