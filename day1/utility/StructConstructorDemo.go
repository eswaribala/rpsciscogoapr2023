package main

import "fmt"

type Trainee struct {
	name       string
	occupation string
	age        int
}

func newTrainee(name string, occupation string, age int) *Trainee {

	p := Trainee{name, occupation, age}
	return &p
}

func main() {

	u := newTrainee("Richard Roe", "driver", 44)

	fmt.Printf("%stores is %d years old and he is a %stores\n", u.name, u.age, u.occupation)
}
