package stores

import (
	"CiscoApr2023/day4/models"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"strconv"
)

func ConnectionHelperORM() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:vignesh@tcp(127.0.0.1:3306)/ciscobankingdb?charset=utf8&parseTime=True")
	//dataSourceName := "root:vignesh@tcp(mysql:3306)/?parseTime=True"
	//db, err := gorm.Open("mysql", dataSourceName)
	db, err := gorm.Open("mysql", "root:vignesh@(mysql:3306)/?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Exec("Create Database ciscobankingdb")
	db.Exec("use ciscobankingdb")
	//generate table
	db.AutoMigrate(&models.Customer{}, &models.Address{})

	return db
}

// SaveCustomer godoc
// @Summary Save a new customer
// @Description Save a new customer with the input payload
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "Create customer"
// @Success 200 {object} models.Customer
// @Router /customers [post]
func SaveCustomer(w http.ResponseWriter, r *http.Request) {

	db := ConnectionHelperORM()
	var customer models.Customer
	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	json.NewDecoder(r.Body).Decode(&customer)
	db.Create(customer)
	tx.Commit()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// GetAllCustomers godoc
// @Summary Get details of all customers
// @Description Get details of all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Customer
// @Router /customers [get]
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	db := ConnectionHelperORM()
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	//db.Find(&customers)
	db.Preloads("Address").Find(&customers)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerById godoc
// @Summary Get details of all customers
// @Description Get details of all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Param accountNo path int true "ID of the Customer"
// @Success 200 {object} models.Customer
// @Router /customers/{accountNo} [get]
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	db := ConnectionHelperORM()
	var customer models.Customer
	params := mux.Vars(r)
	inputAccountNo := params["accountNo"]
	db.First(&customer, inputAccountNo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer godoc
// @Summary Update existing customer
// @Description Update existing customer with the input payload
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "Create customer"
// @Success 200 {object} models.Customer
// @Router /customers [put]
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	db := ConnectionHelperORM()
	var updatedCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	tx := db.Begin()
	db.Save(&updatedCustomer)
	tx.Commit()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Delete customer by id
// @Tags customers
// @Accept  json
// @Produce  json
// @Param accountNo path int true "ID of the Customer"
// @Success 200 {object} models.Customer
// @Router /customers/{accountNo} [delete]
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	db := ConnectionHelperORM()
	params := mux.Vars(r)
	inputAccountNo := params["accountNo"]
	// Convert `traderId` string param to uint64
	id64, _ := strconv.ParseUint(inputAccountNo, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("account_no=?", idToDelete).Delete(&models.Customer{})
	w.WriteHeader(http.StatusNoContent)
}
