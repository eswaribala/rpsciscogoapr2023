package stores

type Transaction struct {
	TransactionId   int64  `json:"transaction_id" gorm:"primary_key"`
	TransactionDate string `json:"transaction_date"`
	Amount          int64  `json:"amount"`
}
