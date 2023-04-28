package test

import (
	models2 "CiscoApr2023/day1/models"
	"CiscoApr2023/day2/models"
	"CiscoApr2023/day2/src"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

var fundtransfer *models.FundTransfer

func init() {
	fundtransfer = new(models.FundTransfer)
	fundtransfer.TransactionId = 1
	fundtransfer.TransactionDate = models2.Date{Day: 28, Month: 4, Year: 2023}
	fundtransfer.TransactionAmount = 5000
}

// file name should have test as lower case
func TestFT(t *testing.T) {
	assert := assert2.New(t)
	assert.NotNil(src.GetFTInstance())

}

func TestGuessFTAmount(t *testing.T) {
	assert := assert2.New(t)

	assert.True(src.GetFTTransactionAmount(fundtransfer) > 10000)

}
