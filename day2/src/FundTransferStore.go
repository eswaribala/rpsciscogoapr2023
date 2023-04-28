package src

import (
	"CiscoApr2023/day2/models"
	"math/rand"
)

func GetFTInstance() *models.FundTransfer {

	return new(models.FundTransfer)
}
func GetFTTransactionAmount(fundTransfer *models.FundTransfer) int64 {

	var fundTransaferInstance *models.FundTransfer
	fundTransaferInstance = fundTransfer
	fundTransaferInstance.TransactionAmount = int64(rand.Int31n(10000)) + fundTransaferInstance.TransactionAmount
	return fundTransaferInstance.TransactionAmount

}
