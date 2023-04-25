package main

import "fmt"

type Member struct {
	name       string
	occupation string
	age        int
}

func main() {

	u := new(Member)
	u.name = "Richard Roe"
	u.occupation = "driver"
	u.age = 44

	fmt.Printf("%stores is %d years old and he is a %stores\n", u.name, u.age, u.occupation)
}
