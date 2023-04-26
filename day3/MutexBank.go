package main

import (
	"fmt"
	"sync"
)

var balance = 50000
var m sync.Mutex
var wgBank sync.WaitGroup

func withdraw(amount int) {
	defer wgBank.Done()
	m.Lock()
	if balance > amount {
		balance = balance - amount
	}
	fmt.Printf("Balance%d\n", balance)
	m.Unlock()
}

func main() {
	fmt.Println("Starting main")
	wgBank.Add(2)
	go withdraw(10000)
	go withdraw(15000)
	wgBank.Wait()
	fmt.Println("Return main")
}
