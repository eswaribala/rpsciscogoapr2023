package main

import (
	"fmt"
	"sync"
)

var balance = 50000
var m sync.Mutex

func withdraw(amount int) {
	m.Lock()
	if balance > amount {
		balance = balance - amount
	}
	fmt.Printf("Balance%d", balance)
	m.Unlock()
}

func main() {
	fmt.Println("Starting main")
	go withdraw(10000)
	go withdraw(15000)
	fmt.Println("Return main")
}
