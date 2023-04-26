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
	db.AutoMigrate(&Customer{}, &Address{})

	return db
}

func SaveCustomer(customer Customer) {

	db := ConnectionHelperORM()

	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	db.Create(customer)
	tx.Commit()

}

func GetAllCustomers() []Customer {
	db := ConnectionHelperORM()
	var customers []Customer
	db.Preloads("Address").Find(&customers)
	return customers
}

func GetCustomerById(accountNo int64) Customer {
	db := ConnectionHelperORM()
	var customer Customer
	db.First(&customer, accountNo)
	return customer
}
func UpdateCustomer(customer Customer) Customer {
	db := ConnectionHelperORM()

	tx := db.Begin()
	db.Save(customer)
	tx.Commit()
	return customer
}

func DeleteCustomer(accountNo int64) {
	db := ConnectionHelperORM()
	db.Where("account_no=?", accountNo).Delete(&Customer{})
}
