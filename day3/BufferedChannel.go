package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func writeCustomerData(pchan chan models.Customer) {

	for i := 0; i < 10; i++ {
		pchan <- models.Customer{
			AccountNo: int64(rand.Int31n(1000000)),
			FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
			Address: models.Address{DoorNo: "13A",
				StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
			DOB: models.Date{
				Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
				Year: 1900 + int(rand.Int31n(122))},
			Email: "Param@gmail.com", Active: true}
	}

	close(pchan)
}

func main() {

	customerBuffCh := make(chan models.Customer, 5)
	go writeCustomerData(customerBuffCh)
	time.Sleep(3 * time.Second)
	for value := range customerBuffCh {
		fmt.Printf("Customer Data=%+v", value)
		time.Sleep(2 * time.Second)
	}
}
