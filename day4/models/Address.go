package models

type Address struct {
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
}
