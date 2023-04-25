package models

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

}

func (transaction *Transaction) Add(date *Date) {

}

func (account *Account) View(printOption *bool, accountNo *int64) (*Account, error) {

}

func (transaction *Transaction) View(printOption *bool, senderAccountNo *int64) (*Transaction, error) {

}
