package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Local variables
	var (
		accountNumber int64
		firstName     string
		//lastName      string
		//address       string
		//email         string
		active bool
	)
	//Reading the values
	fmt.Println("Enter Account No")
	fmt.Scanln(&accountNumber)
	fmt.Println("Enter First Name")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Status")
	fmt.Scanln(&active)

	fmt.Printf("AccountNp%d\n First Name=%s\n Active%t", accountNumber, firstName, active)
	//Type Checking...
	fmt.Println(reflect.TypeOf(accountNumber))
	fmt.Println(reflect.TypeOf(firstName))
	fmt.Println(reflect.TypeOf(active))
}
