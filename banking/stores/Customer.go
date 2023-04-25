package stores

type Customer struct {
	AccountNo int64     `json:"account_no" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   []Address `json:"address" gorm:"ForeignKey:AccountNo"`
	Email     string    `json:"email"`
	DOB       string    `json:"dob"`
	Active    bool      `json:"active"`
}
