package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

func main() {

	customerList := make([]models.Customer, 10)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			customerList[i] = models.Customer{
				AccountNo: int64(rand.Int31n(1000000)),
				FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
				Address: models.Address{DoorNo: "13A",
					StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
				DOB: models.Date{
					Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
					Year: 1900 + int(rand.Int31n(122))},
				Email: "Param@gmail.com", Active: true}
		} else {
			customerList[i] = models.Customer{
				AccountNo: int64(rand.Int31n(1000000)),
				FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
				Address: models.Address{DoorNo: "13A",
					StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
				DOB: models.Date{
					Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
					Year: 1900 + int(rand.Int31n(122))},
				Email: "Param@gmail.com", Active: false}
		}
	}
	sort.SliceStable(customerList, func(i, j int) bool {
		if customerList[i].FirstName == customerList[j].FirstName {
			return customerList[i].LastName < customerList[j].LastName
		} else {
			return customerList[i].FirstName < customerList[j].FirstName
		}

	})

	for _, value := range customerList {
		fmt.Printf("Customer=%+v\n", value)
	}
	//fmt.Printf("First Name=%s", customer.FirstName)
}
