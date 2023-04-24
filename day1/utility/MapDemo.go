package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	customersTable := make(map[models.CustomerType]models.Customer)
	var customerType models.CustomerType
	for i := 0; i < 10; i++ {

		customerType = models.CustomerType{AccountNo: int64(rand.Int31n(1000000)),
			CustomerId: int64(rand.Int31n(1000000))}
		customersTable[customerType] = models.Customer{
			AccountNo: int64(rand.Int31n(1000000)),
			FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
			Address: models.Address{DoorNo: "13A",
				StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
			DOB: models.Date{
				Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
				Year: 1900 + int(rand.Int31n(122))},
			Email: "Param@gmail.com", Active: true}
	}

	keys := make([]models.CustomerType, 10)
	i := 0
	for key, _ := range customersTable {
		fmt.Printf("Key=%d", key)
		keys[i] = key
		i++
	}

}
