package main

import "fmt"

func main() {
	var name string = "Parameswari"
	//pointer
	var pName *string = &name
	fmt.Printf("Name=%s\n", name)
	fmt.Printf("Address%v\n", pName)
	fmt.Printf("Name From Pointer%v\n", *pName)
	//*pName = "Parameswari Bala"
	ChangeName(pName)
	fmt.Printf("Name=%s\n", name)

}

// call by reference
func ChangeName(name *string) {
	*name = "Parameswari Bala"
}
