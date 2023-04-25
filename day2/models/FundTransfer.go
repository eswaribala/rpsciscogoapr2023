package models

import "CiscoApr2023/day1/models"

type FundTransfer struct {
	TransactionId     int64
	TransactionAmount int64
	TransactionDate   models.Date
}
