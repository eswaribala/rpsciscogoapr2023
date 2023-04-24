package main

import (
	"CiscoApr2023/day1/models"
	"encoding/json"
	"fmt"
)

func main() {
	customer := models.Customer{}

	customerJson := `{"account_no": 237867,"first_name":"Param",
         "last_name":"bala","address":{"door_no":"3478",
    "street_name":"gandhi st","city":"Bangalore","state":"KN"},
     "email":"param@gmail.com","dob":{"day":1,"month":5,"year":1988}
    ,"active":true}`
	err := json.Unmarshal([]byte(customerJson), &customer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", customer)

}
