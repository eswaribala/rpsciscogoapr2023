package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg1, wg2 sync.WaitGroup

func readLinkContents(httpResponse *http.Response) {
	defer wg2.Done()
	content, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		errors.New("The content not accessible")
	} else {
		fmt.Println(content)
	}

}

func accessLinkV3(link string) {
	defer wg1.Done()
	wg2.Add(1)
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("Status Code%d of Link%s\n", response.StatusCode, link)
		go readLinkContents(response)
		wg2.Wait()
	}

}

func main() {
	fmt.Println("Kick Starting from main....")
	links := []string{"https://www.google.com", "https://www.rediffmail.com",
		"https://www.facebook.com", "https://pkg.go.dev"}
	wg1.Add(len(links))
	for _, value := range links {
		//delay
		time.Sleep(1 * time.Second)
		go accessLinkV3(value)
	}
	wg1.Wait()

	fmt.Println("Returned to main....")

}
