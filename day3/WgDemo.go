package main

import (
	"fmt"
	"net/http"
	"sync"

	"time"
)

var wg sync.WaitGroup

func accessLinkV2(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("Status Code%d of Link%s\n", response.StatusCode, link)
	}
	wg.Done()
}

func main() {
	fmt.Println("Kick Starting from main....")
	links := []string{"https://www.google.com", "https://www.rediffmail.com",
		"https://www.facebook.com", "https://pkg.go.dev"}
	wg.Add(len(links))
	for _, value := range links {
		//delay
		time.Sleep(1 * time.Second)
		go accessLinkV2(value)
	}
	wg.Wait()

	fmt.Println("Returned to main....")

}
