package models

type RTGS struct {
	FundTransfer FundTransfer
	prevLink     *RTGS
}
