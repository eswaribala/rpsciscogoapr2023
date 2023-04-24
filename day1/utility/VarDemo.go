package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	accountNumber int64
	firstName     string
	lastName      string
	address       string
	email         string
	active        bool
)

func main() {
	args := os.Args
	fmt.Printf("Number of Arguments=%d\n", len(args))
	if len(args) != 7 {
		fmt.Println("Number of args are less than six")
	} else {
		result, err := amendCustomerData(args)
		if err == nil {
			/*for index, value := range result {
				//if index > 0 {
				fmt.Printf("Index=%d,value %s\n", index, value)
				//}
			}*/

			x, _ := strconv.Atoi(result[1])
			accountNumber = int64(x)

			firstName = result[2]
			lastName = result[3]
			address = result[4]
			email = result[5]

		} else {
			fmt.Println(err)
		}
	}

}

func amendCustomerData(data []string) ([]string, error) {

	//dynamically array
	customerData := make([]string, 7)
	accountNo, err := strconv.Atoi(data[1])
	if err == nil {
		if accountNo <= 0 {
			return nil, err
		} else {
			for index, value := range data {
				if index > 0 {

					customerData[index] = value
					if index == 2 {
						customerData[index] = "Mr/Mrs." + value
					}

				} else {

					customerData[0] = strconv.Itoa(int(rand.Int31n(10000)))
				}

			}
			return customerData, nil
		}
	} else {
		return nil, err
	}

}
