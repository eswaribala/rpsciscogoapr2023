package stores

import (
	"database/sql"
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

/*
func GetAllTransactions() {
	db := ConnectionHelper()
	defer db.Close()
	queryString := "select * from transaction"

	rows, err := db.Query(queryString)
	defer rows.Close()
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	var transactions []mockmodels.MTransaction
	for rows.Next() {
		var tObject mockmodels.MTransaction

		err := rows.Scan(&tObject.TransactionId, &tObject.Amount,
			&tObject.Time_Stamp, &tObject.Sender, &tObject.Receiver)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, tObject)
	}

	fmt.Printf("\nTransactions %+v\n", transactions)

}
*/
