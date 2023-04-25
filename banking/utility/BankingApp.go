package main

import (
	"CiscoApr2023/banking/facades"
	"CiscoApr2023/banking/models"
)

func main() {
	//account Instance
	account := models.Account{}
	account.AccountNo = 84388587
	account.RunningTotal = 4365873
	account.OpeningDate = models.Date{Day: 25, Month: 4, Year: 2023}
	account.SavingsAccount.InterestRate = 0.5

	var bankingRef facades.IBank
	bankingRef = &account
	bankingRef.Add()
}
