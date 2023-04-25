package main

import (
	"CiscoApr2023/banking/stores"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	customer := stores.Customer{
		AccountNo: int64(rand.Int31n(1000000)),
		FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
		Address: stores.Address{DoorNo: "13A",
			StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
		DOB:   "23/5/2023",
		Email: "Param@gmail.com", Active: true}
	stores.SaveCustomer(customer)
	for _, value := range stores.GetAllCustomers() {
		fmt.Printf("Customer=%+v", value)
	}
	fmt.Printf("Customer=%+v", stores.GetCustomerById(686045))
	customer.Email = "ciscoparam@gmail.com"
	customer.AccountNo = 686045
	stores.UpdateCustomer(customer)
	stores.DeleteCustomer(243729)
}
