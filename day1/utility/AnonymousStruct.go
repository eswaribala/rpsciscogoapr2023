package main

import "fmt"

func main() {

	u := struct {
		name       string
		occupation string
		age        int
	}{
		name:       "John Doe",
		occupation: "gardener",
		age:        34,
	}

	fmt.Printf("%stores is %d years old and he is a %stores\n", u.name, u.age, u.occupation)
}
