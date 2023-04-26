package main

import (
	"fmt"
	"net/http"
	"time"
)

func accessLinkV1(link string, pChannel chan string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		pChannel <- "link completed the task"
		fmt.Printf("Status Code%d of Link%s\n", response.StatusCode, link)
	}

}

func main() {
	fmt.Println("Kick Starting from main....")
	links := []string{"https://www.google.com", "https://www.rediffmail.com",
		"https://www.facebook.com", "https://pkg.go.dev"}
	//create the channel
	c := make(chan string)
	for _, value := range links {
		//delay
		time.Sleep(1 * time.Second)
		go accessLinkV1(value, c)
	}

	for _, value := range links {

		fmt.Printf("Link%s\t sends message as%s\n", value, <-c)
	}
	fmt.Println("Returned to main....")

}
