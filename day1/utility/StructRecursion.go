package main

import (
	"CiscoApr2023/day1/models"
	"fmt"
)

func main() {
	var tracker models.Tracker // dot notation
	tracker.Id = 1
	tracker.Desc = "Tracking Issues"

	tracker.Issue.Id = 101
	tracker.Issue.Desc = "WIndows Network Issue"

	fmt.Println(tracker)
	fmt.Println("Issue Id:\t", tracker.Issue.Id)

}
