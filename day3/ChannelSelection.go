package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	customerch := make(chan string)
	staffch := make(chan string)
	stockch := make(chan int)

	go func(pStockch chan int) {
		time.Sleep(5 * time.Second)
		pStockch <- int(rand.Int31n(10000))
	}(stockch)
	go customerMessage(customerch)
	go staffMessage(staffch)

	func() {
		select {
		case m1 := <-customerch:
			fmt.Printf("Customer Message%s", m1)
		case m2 := <-staffch:
			fmt.Printf("Staff Message%s", m2)
		case m3 := <-stockch:
			fmt.Printf("Stock Message%d", m3)
		}
	}()

}

func customerMessage(pCustomerch chan string) {
	time.Sleep(4 * time.Second)
	pCustomerch <- "Waiting for Loan Process"
}

func staffMessage(pStaffch chan string) {
	time.Sleep(2 * time.Second)
	pStaffch <- "Process will take 5 to 7 working days"
}
