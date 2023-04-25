package stores

type Address struct {
	AddressId  int64  `json:"address_id" gorm:"primary_key, AUTO_INCREMENT"`
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	AccountNo  int64  `gorm:"column:account_no"`
	Customer   Customer
}
