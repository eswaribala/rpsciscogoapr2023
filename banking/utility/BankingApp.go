package main

import (
	"CiscoApr2023/banking/facades"
	"CiscoApr2023/banking/stores"
	"math/rand"
)

func main() {
	//account Instance
	account := stores.Account{}
	account.AccountNo = int64(rand.Int31n(100000))
	account.RunningTotal = 4365873
	account.OpeningDate = stores.Date{Day: 25, Month: 4, Year: 2023}
	account.SavingsAccount.InterestRate = 0.5

	var bankingRef facades.IBank
	bankingRef = &account
	bankingRef.Add(&stores.Date{Day: 23, Month: 4, Year: 2023})
	bankingRef.View(true)

	transaction := stores.Transaction{}
	transaction.SenderAccountNo = 348768
	transaction.ReceiverAccountNo = 235652
	transaction.Amount = 50000
	transaction.TimeStamp = "11:45:00"
	transaction.Sender = "Parameswari"
	transaction.Receiver = "Vignesh"
	transaction.DirectDebit.PaymentDate = stores.Date{Day: 23, Month: 4, Year: 2023}
	bankingRef = &transaction
	bankingRef.Add(&models.Date{Day: 23, Month: 4, Year: 2023})
	bankingRef.View(false)

}
