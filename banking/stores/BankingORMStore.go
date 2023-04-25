package stores

import "log"

package storeorm

import (
"github.com/amexws/abcbank/mockmodels"
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"log"
)

func ConnectionHelper() *gorm.DB {
	db, err := gorm.Open("mysql", "root:vignesh@tcp(127.0.0.1:3306)/abcormdb?charset=utf8&parseTime=True")

	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	return db
}

func SaveTransaction(transactionId int64, amount int64, timestamp string, sender string, receiver string) {
	mTransaction := mockmodels.MTransaction{TransactionId: transactionId,
		Amount: amount, Time_Stamp: timestamp,
		Sender: sender, Receiver: receiver}
	db := ConnectionHelper()
	//generate table
	db.AutoMigrate(mTransaction);
	tx := db.Begin()

	//To insert or create the record.
	//NOTE: we can insert multiple records too
	db.Debug().Save(mTransaction)
	//Also we can use save that will return primary key
	//resultSave:=db.Debug().Save(newUser)
	tx.Commit()
	//fmt.Println(resultSave)
	//log.Println(resultSave)
}

