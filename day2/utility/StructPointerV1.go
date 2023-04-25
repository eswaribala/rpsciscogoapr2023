package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var customer = new(models.Customer)
	customer.AccountNo = int64(rand.Int31n(1000000))
	customer.FirstName = "Parameswari"
	customer.LastName = "Bala" + strconv.Itoa(int(rand.Int31n(10000)))
	customer.Address = models.Address{DoorNo: "13A",
		StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"}
	customer.DOB = models.Date{
		Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
		Year: 1900 + int(rand.Int31n(122))}
	customer.Email = "Param@gmail.com"
	customer.Active = true
	changeCustomerNameV1(customer)
	fmt.Printf("Customer=%+v", customer)
}

func changeCustomerNameV1(customer *models.Customer) {
	customer.FirstName = "Vignesh"
}
