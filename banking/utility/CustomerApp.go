package main

import (
	"CiscoApr2023/banking/stores"
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	customer := models.Customer{
		AccountNo: int64(rand.Int31n(1000000)),
		FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
		Address: models.Address{DoorNo: "13A",
			StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
		DOB: models.Date{
			Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
			Year: 1900 + int(rand.Int31n(122))},
		Email: "Param@gmail.com", Active: true}
	stores.SaveCustomer(customer)
	for _, value := range stores.GetAllCustomers() {
		fmt.Printf("Customer=%+v", value)
	}
	//fmt.Printf("Customer=%+v", stores.GetCustomerById(64852))
	//customer.Address.City = "Pune"
	//stores.UpdateCustomer(customer)

}
