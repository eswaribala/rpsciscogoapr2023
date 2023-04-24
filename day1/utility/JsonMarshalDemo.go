package main

import (
	"CiscoApr2023/day1/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var customer models.Customer
	customer = models.Customer{
		AccountNo: int64(rand.Int31n(1000000)),
		FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
		Address: models.Address{DoorNo: "13A",
			StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
		DOB: models.Date{
			Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
			Year: 1900 + int(rand.Int31n(122))},
		Email: "Param@gmail.com", Active: true}

	customerJson, _ := json.Marshal(customer)
	fmt.Printf("Json Customer%s", customerJson)

}
