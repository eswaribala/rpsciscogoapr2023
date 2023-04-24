package main

import (
	"fmt"
	"sort"
)

func main() {

	names := []string{"Param", "Vignesh", "Bala"}
	sort.Strings(names)
	fmt.Println(names)
	scores := []int{67, 15, 56, 98, 25}
	sort.Ints(scores)
	fmt.Println(scores)
}
