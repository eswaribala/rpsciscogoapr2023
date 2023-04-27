package stores

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func ConnectionHelperORM() *gorm.DB {
	db, err := gorm.Open("mysql", "root:vignesh@tcp(127.0.0.1:3306)/ciscobankingdb?charset=utf8&parseTime=True")

	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	//generate table
	db.AutoMigrate(&Transaction{})

	return db
}

func SaveTransaction(transaction *Transaction) {

	db := ConnectionHelperORM()

	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	db.Create(transaction)
	tx.Commit()

}

func GetAllTransactions() ([]*Transaction, error) {
	db := ConnectionHelperORM()
	var resultArr []*Transaction
	db.Find(&resultArr)

	return resultArr, nil
}
