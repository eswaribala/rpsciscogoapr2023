package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	banks := make([]string, 5, 8)
	for i := 0; i < 10; i++ {
		banks = append(banks, "Bank"+strconv.Itoa(int(rand.Int31n(10000))))
	}
	fmt.Println(banks)
	//banks = banks[4:6]
	//banks = append(banks, "HSBC")
	//banks = append(banks, "HDFC")
	fmt.Println(banks[5:8])
}
