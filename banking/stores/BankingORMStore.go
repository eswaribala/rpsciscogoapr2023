package stores

import (
	"CiscoApr2023/day1/models"
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
	return db
}

func SaveCustomer(customer models.Customer) {

	db := ConnectionHelperORM()
	//generate table
	db.AutoMigrate(customer)
	tx := db.Begin()
	db.Save(customer)
	tx.Commit()

}

func GetAllCustomers() []models.Customer {
	db := ConnectionHelperORM()
	var customers []models.Customer
	db.Find(&customers)
	return customers
}

func GetCustomerById(accountNo int64) models.Customer {
	db := ConnectionHelperORM()
	var customer models.Customer
	db.First(&customer, accountNo)
	return customer
}
func UpdateCustomer(customer models.Customer) models.Customer {
	db := ConnectionHelperORM()

	tx := db.Begin()
	db.Save(customer)
	tx.Commit()
	return customer
}

func DeleteCustomer(accountNo int64) {
	db := ConnectionHelperORM()
	db.Where("account_no=?", accountNo).Delete(&models.Customer{})
}
