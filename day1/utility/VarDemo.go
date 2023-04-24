package main

import (
	"fmt"
	"os"
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
			for index, value := range result {
				if index > 0 {
					fmt.Printf("Index=%d,value %s\n", index, value)
				}
			}
		} else {
			fmt.Println(err)
		}
	}

}

func amendCustomerData(data []string) (updatedData []string, error) {

}
