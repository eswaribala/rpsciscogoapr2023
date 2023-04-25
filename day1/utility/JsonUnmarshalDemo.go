package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	users, _ := ioutil.ReadFile("users.json")
	var result []map[string]interface{}

	json.Unmarshal([]byte(users), &result)

	for outerKey, outerValue := range result {

		if reflect.TypeOf(outerValue).String() == "map[string]interface {}" {
			for innerKey, innerValue := range outerValue {

				if reflect.TypeOf(innerValue).String() == "map[string]interface {}" {
					fmt.Printf("%+v", innerValue)

				} else {
					fmt.Printf("key=%stores,value%stores", innerKey, innerValue)
				}
			}
		} else {
			fmt.Printf("key=%stores,value%stores", outerKey, outerValue)
		}

	}

}
