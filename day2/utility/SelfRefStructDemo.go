package main

import (
	models2 "CiscoApr2023/day1/models"
	"CiscoApr2023/day2/models"
	"math/rand"
)

func main() {
	var rtgsList []*models.RTGS
	for i := 0; i < 10; i++ {
		rtgsList[i] = new(models.RTGS)
		rtgsList[i].FundTransfer = models.FundTransfer{
			TransactionId:     int64(rand.Int31n(1000000)),
			TransactionAmount: int64(rand.Int31n(1000000)),
			TransactionDate: models2.Date{
				Day:   1 + int(rand.Int31n(26)),
				Month: 1 + int(rand.Int31n(11)),
				Year:  1900 + int(rand.Int31n(122))}}
		if i > 0 {
			rtgsList[i].Link = rtgsList[i-1]
		}
	}

}
