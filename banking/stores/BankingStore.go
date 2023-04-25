package stores

import (
	"CiscoApr2023/banking/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectionHelper() *sql.DB {
	db, err := sql.Open("mysql", "root:vignesh@(127.0.0.1:3306)/ciscobankingdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateAccount(accountNo int64, runningTotal int64, openingDate string, roi float32) (int64, error) {

	db := ConnectionHelper()
	defer db.Close()
	queryString := "Insert into savingsaccount (Account_No,Running_Total,Opening_Date,Interest_Rate) values(?,?,?,?)"
	result, err := db.Exec(queryString, accountNo, runningTotal, openingDate, roi)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	return result.RowsAffected()
}

func GetAllAccounts() {
	db := ConnectionHelper()
	defer db.Close()
	queryString := "select * from savings_account"

	rows, err := db.Query(queryString)
	defer rows.Close()
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	var accounts []models.Account
	for rows.Next() {
		var tObject models.Account

		err := rows.Scan(&tObject.AccountNo, &tObject.RunningTotal,
			&tObject.OpeningDate, &tObject.SavingsAccount.InterestRate)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, tObject)
	}

	fmt.Printf("\nTransactions %+v\n", accounts)

}
