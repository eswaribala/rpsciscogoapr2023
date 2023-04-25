package models

type RTGS struct {
	FundTransfer FundTransfer
	Link         *RTGS
}
