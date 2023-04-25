package stores

import (
	"fmt"
)

type Account struct {
	AccountNo      int64
	RunningTotal   int64
	OpeningDate    string
	SavingsAccount struct {
		InterestRate float32
	}
	CurrentAccount struct {
		ODLimit int64
	}
}

type DirectDebit struct {
	PaymentDate string
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

func (account *Account) Add(date *string) {

	fmt.Printf("Account=%+v\n", account)
	//dateString := strconv.Itoa(account.OpeningDate.Day) + "/" + strconv.Itoa(account.OpeningDate.Month) + "/" + strconv.Itoa(account.OpeningDate.Year)
	count, err := CreateAccount(account.AccountNo,
		account.RunningTotal, account.OpeningDate, account.SavingsAccount.InterestRate)
	if err == nil {
		fmt.Printf("Account Created=%d\n", count)
	} else {
		fmt.Println(err)
	}
}

func (account *Account) View(permission bool) {
	GetAllAccounts()
	if permission {
		fmt.Println("Print the Receipt")
	} else {
		fmt.Println("Go Green Thank you")
	}

}

func (account *Account) ViewById(id *int64) {

}
func (account *Account) Update(runningTotal *int64) {

}
func (account *Account) Delete(id *int64) {

}
func (transaction *Transaction) Add(date *string) {
	fmt.Printf("Transaction=%+v\n", transaction)
}
func (transaction *Transaction) View(permission bool) {
	if permission {
		fmt.Println("Print the Receipt")
	} else {
		fmt.Println("Go Green Thank you")
	}
}

func (transaction *Transaction) ViewById(id *int64) {

}
func (transaction *Transaction) Update(runningTotal *int64) {

}
func (transaction *Transaction) Delete(id *int64) {

}
