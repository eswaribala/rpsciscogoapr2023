package models

import (
	"fmt"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

type Account struct {
	AccountNo      int64
	RunningTotal   int64
	OpeningDate    Date
	SavingsAccount struct {
		InterestRate float32
	}
	CurrentAccount struct {
		ODLimit int64
	}
}

type DirectDebit struct {
	PaymentDate Date
}

type Transaction struct {
	SenderAccountNo   int64
	ReceiverAccountNo int64
	Amount            int64
	Sender            string
	Receiver          string
	TimeStamp         string
	DirectDebit       DirectDebit
}

func (account *Account) Add(date *Date) {

	fmt.Printf("Account=%+v", account)
}

func (transaction *Transaction) Add(date *Date) {
	fmt.Printf("Transaction=%+v", transaction)
}

func (account *Account) View(permission bool) {

}

func (transaction *Transaction) View(permission bool) {

}
