package stores

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectionHelper() *sql.DB {
	db, err := sql.Open("mysql", "root:vignesh@(127.0.0.1:3306)/abcbankdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTransaction(transactionId int64, amount int64, timestamp string, sender string, receiver string) (int64, error) {

	db := ConnectionHelper()
	defer db.Close()
	queryString := "Insert into transaction (Amount,Time_Stamp,Sender,Receiver) values(?,?,?,?)"
	result, err := db.Exec(queryString, amount, timestamp, sender, receiver)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	return result.RowsAffected()
}

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
