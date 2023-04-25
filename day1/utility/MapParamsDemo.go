package main

import "fmt"

func main() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	Display(student)
}

func Display(student map[string]string) {
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}
