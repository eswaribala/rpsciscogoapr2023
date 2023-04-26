package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

/*
Quick Read Whole File to Memory
*/
func main() {

	data, _ := ioutil.ReadFile("employeedata.csv")
	r := csv.NewReader(strings.NewReader(string(data)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
