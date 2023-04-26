package main

import (
	"fmt"
	"net/http"
)

func accessLink(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Status Code%d of Link%s\n", response.StatusCode, link)
	}

}

func main() {
	fmt.Println("Kick Starting from main....")
	links := []string{"https://www.google.com", "https://www.rediffmail.com",
		"https://www.facebook.com", "https://pkg.go.dev"}
	for _, value := range links {
		accessLink(value)
	}
	fmt.Println("Returned to main....")

}
