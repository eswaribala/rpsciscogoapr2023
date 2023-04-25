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

	fmt.Printf("%s is %d years old and he is a %s\n", u.name, u.age, u.occupation)
}
