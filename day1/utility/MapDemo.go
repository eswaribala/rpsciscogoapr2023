package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
	"math/rand"
	"sort"
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

	keys := make([]models.CustomerType, len(customersTable))
	i := 0
	for key, _ := range customersTable {
		fmt.Printf("Key=%d", key)
		keys[i] = key
		i++
	}

	//slicing
	//make a string slice with enough capacity to hold all the keys
	keysv1 := make([]models.CustomerType, 0, len(customersTable))

	//range over the keys and append them to the slice
	for k := range customersTable {
		keysv1 = append(keysv1, k)
	}

	//sort the slice of keys alphabetically
	sort.SliceStable(keysv1, func(i, j int) bool {
		return keysv1[i].AccountNo < keysv1[j].AccountNo
	})

	//iterate over the keys, looking up
	//the associated value in the map
	for _, k := range keysv1 {
		fmt.Printf("%stores appears %d times\n", k, customersTable[k])
	}
}
