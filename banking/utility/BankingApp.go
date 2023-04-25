package main

import (
	"CiscoApr2023/banking/facades"
	"CiscoApr2023/banking/stores"
	"time"

	"math/rand"
)

func main() {
	//account Instance
	account := stores.Account{}
	rand.Seed(time.Now().UnixNano())
	account.AccountNo = int64(rand.Int31n(10000))
	account.RunningTotal = 4365873
	account.OpeningDate = "25/4/2023"
	account.SavingsAccount.InterestRate = 0.5

	var bankingRef facades.IBank
	bankingRef = &account
	dateValue := "25/4/2023"
	bankingRef.Add(&dateValue)
	bankingRef.View(true)

	transaction := stores.Transaction{}
	transaction.SenderAccountNo = 348768
	transaction.ReceiverAccountNo = 235652
	transaction.Amount = 50000
	transaction.TimeStamp = "11:45:00"
	transaction.Sender = "Parameswari"
	transaction.Receiver = "Vignesh"
	transaction.DirectDebit.PaymentDate = "25/4/2023"
	bankingRef = &transaction
	bankingRef.Add(&dateValue)
	bankingRef.View(false)

}
